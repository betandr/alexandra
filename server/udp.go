package server

import (
	"fmt"
	"net"
)

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("Hello, Client!"), addr)
	if err != nil {
		fmt.Printf("error sending response: %v", err)
	}
}

// Start the UDP server, listening on port 8067
func Start() {
	p := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: 8067,
		IP:   net.ParseIP("127.0.0.1"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	for {
		_, remoteaddr, err := ser.ReadFromUDP(p)
		fmt.Printf("message from %v: %s \n", remoteaddr, p)
		if err != nil {
			fmt.Printf("error reading: %v", err)
			continue
		}
		go sendResponse(ser, remoteaddr)
	}
}
