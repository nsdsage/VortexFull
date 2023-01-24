
package main

import (
    _ "encoding/json"
    _ "context"
    "fmt"
    "log"
    _ "math"
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "bdns/api"
)

type Domain struct {
	IpAddress  string
	DeviceType string
}

var client *ethclient.Client
var err error

var records = map[string]string{
	"mail.amazon.com":  "192.162.1.2",
	"paste.amazon.com": "191.165.0.3",
}

func lookupFunc(lookup string, seclvl int64) (string, error) {
	//Do some action
	//Get data from DB
	//Process it further more
	// return "192.2.2.1", nil
	//var ips []string
	client, err = ethclient.Dial("http://192.168.1.6:8545")
        if err != nil {
                fmt.Println("Failed at Dial")
                log.Fatal(err)
        }
	conn, err := api.NewApi(common.HexToAddress("0xC6F133979422d9e19B562aBa82D540c86FcD1747"), client)
        if err != nil {
                fmt.Println(err)
        }
	domains, _ := conn.GetRecord(&bind.CallOpts{}, lookup, big.NewInt(seclvl))
	for _, domain := range domains {
		return domain.IpAddress, nil
	}
	return "", nil
}
func main() {
	client, err = ethclient.Dial("http://192.168.1.6:8545")
	if err != nil {
		fmt.Println("Failed at Dial")
        	log.Fatal(err)
    	}
	_ = client
	var googleRecords = map[string]string{
		"mail.google.com":  "192.168.0.2",
		"paste.google.com": "192.168.0.3",
		"google.com": "40.76.4.15",
	}
	var microsoftRecords = map[string]string{
		"mail.microsoft.com":  "192.168.0.78",
		"paste.microsoft.com": "192.168.0.25",
	}
//	var vortexRecords = map[string]string{
//		"mail.vortex.mil": "127.0.0.12",
//	}
	dns := NewDNSServer(53, 853)
	// tcpdns := NewDNSServer(853, "tcp")
	dns.AddZoneData("google.com", googleRecords, nil, DNSForwardLookupZone)
	dns.AddZoneData("microsoft.com", microsoftRecords, nil, DNSForwardLookupZone)
	dns.AddZoneData("vortex.mil", nil, lookupFunc, DNSForwardLookupZone)

	/* Incase if the records are not static or to be taken from DB or from any other sources
	lookupFunc method can be used.append*/
	//dns.AddZoneData("amazon.com", nil, lookupFunc, DNSForwardLookupZone)
	dns.StartAndServe()
}

