package main

import (
        "fmt"
        "log"

        "github.com/surgemq/surgemq/service"
)

var schedule = 60

func main() {
        fmt.Println("Starting")
        startServer()
}

func startServer() {
        svr := &service.Server{
                KeepAlive:        300,               // seconds
                ConnectTimeout:   2,                 // seconds
                SessionsProvider: "mem",             // keeps sessions in memory
                Authenticator:    "mockSuccess",     // always succeed
                TopicsProvider:   "mem",             // keeps topic subscriptions in memory
        }
        if err := svr.ListenAndServe("tcp://:8080"); err != nil {
                log.Fatal(err)
	}
}
