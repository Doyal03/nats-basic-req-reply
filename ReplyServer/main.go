package main

import (
	"fmt"
	"runtime"

	"github.com/nats-io/nats.go"
)

func main() {
	conn, err := nats.Connect(nats.DefaultURL) // localhost:4222
	if err != nil {
		panic(err)
	}
	for conn.IsConnected() {
		conn.Subscribe("foo", func(msg *nats.Msg) {
			fmt.Printf("%v\n", string(msg.Data))
			conn.Publish(msg.Reply, []byte("Message delivered"))
		})
		runtime.Goexit()
	}
	defer conn.Close()
}
