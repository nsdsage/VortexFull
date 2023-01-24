package main

import (
        "fmt"
//        "time"

        emitter "github.com/emitter-io/go"
)

var key = "pPhQSVsGmqjJQFEaggxKzXxYdbsICAoc"
var channel = "sync/test"

func main() {
        o := emitter.NewClientOptions()
        o.SetOnMessageHandler(func(client emitter.Emitter, msg emitter.Message) {
                fmt.Printf("Received message: %s\n", msg.Payload())
        })
        o.AddBroker("tcp://10.10.100.246:9089")

        c := emitter.NewClient(o)
        sToken := c.Connect()
        if sToken.Wait() && sToken.Error() != nil {
                panic("Error on Client.Connect(): " + sToken.Error().Error())
        }
        fmt.Println(o)

        c.SubscribeWithHistory(key, "sync", 1)
        for { }
}

