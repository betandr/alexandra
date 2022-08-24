package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"rtp"
)

func main() {
	addr := net.UDPAddr{
		Port: 8067,
		IP:   net.ParseIP("127.0.0.1"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("error listening: %s", err)
		return
	}

	fmt.Println("starting...")
	for {
		var p rtp.Packet
		dec := gob.NewDecoder(conn)
		err := dec.Decode(&p)
		if err != nil {
			log.Println("decoding: ", err)
		}

		fmt.Printf("packet payload recieved: %s \n", string(p.Payload))
	}
}
