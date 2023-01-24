package main

import (
        "fmt"
        "log"
        "time"
        "strconv"
        "encoding/json"
        "github.com/surgemq/message"
        "github.com/surgemq/surgemq/service"
)

var schedule = 60
var stopped = true
var latest = ""
var state = false;

type Payload_msg struct {
	Latest string
	State bool
}

func main() {
    fmt.Println("Starting")
    c := &service.Client{}
	msg := message.NewConnectMessage()
	msg.SetVersion(4)
	msg.SetWillTopic([]byte("sync"))
	msg.SetCleanSession(true)
	msg.SetKeepAlive(300)
	if err := c.Connect("tcp://127.0.0.1:3011", msg); err != nil {
		log.Fatal(err)
	}
	go subscribe_msg(c)
	heartbeat(c)
	fmt.Println("FIN")
}
func subscribe_msg(c *service.Client) {
	submsg := message.NewSubscribeMessage()
	submsg.AddTopic([]byte("#"), 0)
	c.Subscribe(submsg, nil, func(msg *message.PublishMessage) error {
		return nil
	})
}

func heartbeat(c *service.Client) {
	for {
		msg := message.NewPublishMessage()
		now := time.Now()
		next := time.Minute * time.Duration(1)
		beat := now.Add(next)
		state = !state;
		latest = strconv.FormatInt(beat.UnixNano(), 10)
		pyld_msg := &Payload_msg{Latest: latest, State: state}
		rsp, err := json.Marshal(pyld_msg)
		if (err == nil) {
			msg.SetPayload([]byte(rsp))
		}
		msg.SetTopic([]byte("sync"))
		msg.SetRetain(true)
		fmt.Println(latest)
		c.Publish(msg, nil)
		time.Sleep(time.Duration(schedule) * time.Second)
	}
}
