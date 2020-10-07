package udp

import (
	"fmt"
	"log"
	"net"
)

func respond(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("Hello back, Client!"), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

// Listen on UDP port 8067
func Listen() error {
	p := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: 8067,
		IP:   net.ParseIP("127.0.0.1"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		return err
	}
	for {
		_, remoteaddr, err := ser.ReadFromUDP(p)
		log.Printf("message from %v %s \n", remoteaddr, p)
		if err != nil {
			fmt.Printf("error reading from UDP:  %v", err)
			continue
		}
		go respond(ser, remoteaddr)
	}
}
