package main

import (
	"fmt"
	"net"
)

func main() {
	addr := net.UDPAddr{
		Port: 8067,
		IP:   net.ParseIP("127.0.0.1"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("error listening: %s", err)
		return
	}

	buf := make([]byte, 2048)

	fmt.Println("starting...")
	for {
		_, remoteaddr, err := ser.ReadFromUDP(buf)
		fmt.Printf("message from %s: %s \n", remoteaddr, buf)
		if err != nil {
			fmt.Printf("error reading: %v", err)
			continue
		}

		_, err = ser.WriteToUDP([]byte("Hello, Client!"), remoteaddr)
		if err != nil {
			fmt.Printf("error sending response: %v", err)
		}
	}
	fmt.Println("done...")
}
