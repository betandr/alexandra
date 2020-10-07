package main

import (
	"log"

	"github.com/betandr/alexandra/rtp"
	"github.com/betandr/alexandra/udp"
)

func main() {
	go udp.Listen()

	packet := rtp.Packet{
		Header:  rtp.NewHeader(),
		Payload: "Hello, Server!",
	}

	var c udp.Client
	c.Connect("127.0.0.1:8067")
	response, err := c.Send(packet)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("from server: %v\n", response)
}
