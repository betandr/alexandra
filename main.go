package main

import (
	"fmt"
	"os"

	"github.com/betandr/alexandra/client"
	"github.com/betandr/alexandra/rtp"
	"github.com/betandr/alexandra/server"
)

func main() {
	go server.Start()

	conn, err := client.NewConnection()
	if err != nil {
		fmt.Printf("error getting connection: %v\n", err)
		os.Exit(1)
	}

	packet := rtp.Packet{
		Header:  rtp.NewHeader(),
		Payload: "Hello, World!",
	}

	client.Send(conn, packet)
}
