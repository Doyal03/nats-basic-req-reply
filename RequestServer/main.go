package main

import (
	// "fmt"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	conn, err := nats.Connect(nats.DefaultURL) // localhost:4222
	if err != nil {
		panic(err)
	}
	if conn.IsConnected() { // will print true if server is up
		for i := 5; i >= 1; i-- {
			_, err := conn.Request("foo", []byte(fmt.Sprint(i)), 1000*time.Millisecond)
			if err != nil {
				panic(err)
			}
			time.Sleep(500 * (time.Millisecond))
		}
		msg, err := conn.Request("foo", []byte("BOOM!"), 1000*time.Millisecond)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(msg.Data))
	}
	defer conn.Close()
}
