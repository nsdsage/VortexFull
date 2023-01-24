package main

import (
	"os"
	"fmt"
	"net"
	"bufio"
	"strconv"
	"encoding/json"

	"github.com/golang-jwt/jwt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// VARS & STRUCTS
type ZoneType uint16
type customHandler func(string, int64) (string, error)

type dnsClaim struct {
	Credentials struct {
		User  int    `json:"User"`
		Token string `json:"Token"`
	} `json:"Credentials"`
	Request struct {
		Type    string `json:"Type"`
		Version string `json:"Version"`
		Address string `json:"Address"`
		Sec int64 `json:"Sec"`
	} `json:"Request"`
	jwt.StandardClaims
}

// UDP Connection Struct
type udpConnection struct {
	conn net.PacketConn
	addr net.Addr
}
func (udp *udpConnection) Write(b []byte) error {
	udp.conn.WriteTo(b, udp.addr)
	return nil
}

const (
	DNSForwardLookupZone ZoneType = 1
	DNSReverseLookupZone ZoneType = 2 //Todo
)
type Handler interface {
	serveDNS(*udpConnection, *layers.DNS, net.Conn, string, string, int64)
}
type serveMux struct {
	handler map[string]Handler
}
type jsonLookup struct {
	Name string
	Address string
}
//DNSServer is the contains the runtime information
type DNSServer struct {
	port	int
	tcpport	int
	handler Handler
	protocol string
}

//NewDNSServer - Creates new DNSServer //////////////////////
func NewDNSServer(port int, tcpport int) *DNSServer {
	handler := NewServeMux()
	return &DNSServer{port: port, tcpport: tcpport, handler: handler}
}
//NewServeMux - creates a servemux with handler field intialised
func NewServeMux() *serveMux {
	h := make(map[string]Handler)
	return &serveMux{handler: h}
}
/////////////////////////////////////////////////////////////

//// GENERATE NEW RECORD FUNCTIONS //////////////////////////
//AddZoneData - Depending on the zoneType and recordType  this function generates appropriate handler and registers in the serveMux
func (dns *DNSServer) AddZoneData(zone string, records map[string]string, lookupFunc func(string, int64) (string, error), lookupZone ZoneType) {
	if lookupZone == DNSForwardLookupZone {
		serveMuxCurrent := dns.handler.(*serveMux)
		serveMuxCurrent.handleFunc(zone, generateHandler(records, lookupFunc))
	}
}
func generateHandler(records map[string]string, lookupFunc customHandler) func(w *udpConnection, r *layers.DNS, conn net.Conn, domain string, protocol string, seclvl int64){
	return func(w *udpConnection, r *layers.DNS, conn net.Conn, domain string, protocol string, seclvl int64) {
		// switch r.Questions[0].Type {
		// case layers.DNSTypeA:
			handleATypeQuery(w, r, records, lookupFunc, conn, domain, protocol, seclvl)
		// }
	}
}
func (srv *serveMux) handleFunc(pattern string, f func(*udpConnection, *layers.DNS, net.Conn, string, string, int64)) {
	srv.handler[pattern] = handlerConvert(f)
}
/////////////////////////////////////////////////////////////

// Start Server and Listen for UDP/TCP Connections //////////
func (dns *DNSServer) StartAndServe() {
	go dns.StartAndServeUDP()
	dns.StartAndServeTCP()
}
func (dns *DNSServer) StartAndServeTCP() {
	listener, err := net.Listen("tcp" , "0.0.0.0:"+strconv.Itoa(dns.tcpport))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Listening on with " +"tcp 127.0.0.1:" + strconv.Itoa(dns.tcpport))
	for {
		// Listen for an incoming connection.
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		dns.serveTCP(conn)
		fmt.Println("Accepted Connection")
	}
}
func (dns *DNSServer) StartAndServeUDP() {
	addr := net.UDPAddr{
		Port: dns.port,
		IP:   net.ParseIP("0.0.0.0"),
	}
	listener, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()
	udpConnection := &udpConnection{conn: listener}
	fmt.Println("Listening on with " +dns.protocol+" 127.0.0.1:" + strconv.Itoa(dns.port))
	dns.serveUDP(udpConnection)
}

func (dns *DNSServer) serveUDP(u *udpConnection) {
	for {
		buf := make([]byte, 1024)
		_, addr, _ := u.conn.ReadFrom(buf)
		u.addr = addr
		packet := gopacket.NewPacket(buf, layers.LayerTypeDNS, gopacket.Default)
		dnsPacket := packet.Layer(layers.LayerTypeDNS)
		tcp, _ := dnsPacket.(*layers.DNS)
		dns.handler.serveDNS(u, tcp, nil, addr.String(), "udp", 0)
	}
}
func (dns *DNSServer) serveTCP(conn net.Conn) {
	tokenString, _ := bufio.NewReader(conn).ReadString('\n')
	token, err := jwt.ParseWithClaims(tokenString, &dnsClaim{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("$C&F)J@NcRfUjXn2r5u8x/A%D*G-KaPd"), nil
	})
	if claims, ok := token.Claims.(*dnsClaim); ok && token.Valid {
		seclvl := claims.Request.Sec
		dns.handler.serveDNS(nil, nil, conn, claims.Request.Address, "tcp", seclvl)
	} else {
			fmt.Println(err)
	}
//	conn.Close()
}
type handlerConvert func(*udpConnection, *layers.DNS, net.Conn, string, string, int64)

func (f handlerConvert) serveDNS(w *udpConnection, r *layers.DNS, conn net.Conn, domain string, procotol string, seclvl int64) {
 	f(w, r, conn, domain, procotol, seclvl)
}
func (srv *serveMux) serveDNS(u *udpConnection, request *layers.DNS, conn net.Conn, domain string, procotol string, seclvl int64) {
	if (procotol == "tcp") {
		var h Handler
		if h = srv.match(string(domain), layers.DNSTypeA); h == nil {
			//todo: log handler not found
			fmt.Println("no handler found for ", string(domain))
		} else {
			h.serveDNS(nil, nil, conn, domain, "tcp", seclvl)
		}
	} else {
		var h Handler
		if len(request.Questions) < 1 { // allow more than one question
			return
		}
	   fmt.Println(string(request.Questions[0].Type))
		if h = srv.match(string(request.Questions[0].Name), request.Questions[0].Type); h == nil {
			//todo: log handler not found
			fmt.Println("no handler found for ", string(request.Questions[0].Name))
		} else {
			h.serveDNS(u, request, nil, "", "udp", seclvl)
	   }
	}
}

/////////////////////////////////////////////////////////////
// Core DNS Internals

func (srv *serveMux) match(q string, t layers.DNSType) Handler {
	var handler Handler
	b := make([]byte, len(q)) // worst case, one label of length q
	off := 0
	end := false
	for {
		l := len(q[off:])
		for i := 0; i < l; i++ {
			b[i] = q[off+i]
			if b[i] >= 'A' && b[i] <= 'Z' {
				b[i] |= 'a' - 'A'
			}
		}
		if h, ok := srv.handler[string(b[:l])]; ok { // causes garbage, might want to change the map key
			if uint16(t) != uint16(43) {
				return h
			}
			// Continue for DS to see if we have a parent too, if so delegeate to the parent
			handler = h
		}
		off, end = nextLabel(q, off)
		if end {
			break
		}
	}
	// Wildcard match, if we have found nothing try the root zone as a last resort.
	if h, ok := srv.handler["."]; ok {
		return h
	}
	return handler
}

func nextLabel(s string, offset int) (i int, end bool) {
	quote := false
	for i = offset; i < len(s)-1; i++ {
		switch s[i] {
		case '\\':
			quote = !quote
		default:
			quote = false
		case '.':
			if quote {
				quote = !quote
				continue
			}
			return i + 1, false
		}
	}
	return i + 1, true
}

func handleATypeQuery(w *udpConnection, r *layers.DNS, records map[string]string, lookupFunc customHandler, conn net.Conn, domain string, protocol string, seclvl int64) {
	var dnsAnswer layers.DNSResourceRecord
	dnsAnswer.Type = layers.DNSTypeA
	var ip string
	var err error
	var ok bool
	if lookupFunc == nil {
		if (protocol == "udp") {
			ip, ok = records[string(r.Questions[0].Name)]
		} else {
			ip, ok = records[domain]
		}
		if !ok {
			//Todo: Log no data present for the IP and handle:todo
		}
	} else {
		if (protocol == "udp") {
			ip, err = lookupFunc(string(r.Questions[0].Name), seclvl)
		} else {
			ip, err = lookupFunc(domain, seclvl)
		}
	}
	a, _, _ := net.ParseCIDR(ip + "/24")
	dnsAnswer.Type = layers.DNSTypeA
	dnsAnswer.IP = a
	if (protocol == "udp") {
		dnsAnswer.Name = []byte(r.Questions[0].Name)
	} else {
		dnsAnswer.Name = []byte(domain)
	}
	dnsAnswer.Class = layers.DNSClassIN
	
	replyMess := r
	if (replyMess != nil) {
		replyMess.QR = true
		replyMess.ANCount = 1
		// replyMess.OpCode = layers.DNSOpCodeNotify
		replyMess.OpCode = layers.DNSOpCodeQuery
		replyMess.AA = true
		replyMess.Answers = append(replyMess.Answers, dnsAnswer)
		replyMess.ResponseCode = layers.DNSResponseCodeNoErr
		buf := gopacket.NewSerializeBuffer()
		opts := gopacket.SerializeOptions{} // See SerializeOptions for more details.
		err = replyMess.SerializeTo(buf, opts)
		if err != nil {
			fmt.Println(err)
		}
		w.Write(buf.Bytes())
	} else {
		jv := &jsonLookup{domain, dnsAnswer.IP.String()}
		jsonValue, err := json.Marshal(jv)
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.Write([]byte(jsonValue))
	}
}
