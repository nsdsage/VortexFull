package main

import (
        "time"
        "strconv"
        emitter "github.com/emitter-io/go"
)

var key = "pPhQSVsGmqjJQFEaggxKzXxYdbsICAoc"
var channel = "sync/timing"
var beating = true
var schedule = 60

func main() {
        o := emitter.NewClientOptions()
        o.AddBroker("tcp://127.0.0.1:8080")
        c := emitter.NewClient(o)
        sToken := c.Connect()
        if sToken.Wait() && sToken.Error() != nil {
                panic("Error on Client.Connect(): " + sToken.Error().Error())
                beating = false
        }
        heartbeat(c)
        c.Disconnect(0)
}

func heartbeat(c emitter.Emitter) {
        for (beating) {
                token := c.PublishWithTTL(key, channel, strconv.FormatInt(time.Now().UnixNano(), 10), 60)
                token.Wait()
                time.Sleep(time.Duration(schedule) * time.Second)
        }
}
