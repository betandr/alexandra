package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"rtp"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8067")
	if err != nil {
		fmt.Printf("error dialling: %s", err)
		return
	}
	defer conn.Close()

	packet := rtp.Packet{
		Header:  rtp.NewHeader(),
		Payload: []byte("Hello, UDP!"),
	}

	gob.Register(rtp.Header{})
	enc := gob.NewEncoder(conn)
	err = enc.Encode(packet)
	if err != nil {
		log.Fatal("encode: ", err)
	}
}
