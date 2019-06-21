package client

import (
	"bufio"
	"fmt"
	"net"
)

// Start the UDP client, sending a message on port 8067
func Start() {
	conn, err := net.Dial("udp", "127.0.0.1:8067")
	defer conn.Close()
	if err != nil {
		fmt.Printf("error dialing udp: %v", err)
		return
	}

	p := make([]byte, 2048)
	fmt.Fprintf(conn, "Hello, UDP Server!")
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		fmt.Printf("server responded: %s\n", p)
	} else {
		fmt.Printf("error reading: %v\n", err)
	}
}
