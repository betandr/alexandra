package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8067")
	if err != nil {
		fmt.Printf("error dialling: %s", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 2048)
	fmt.Fprintf(conn, "Hello, UDP!")

	_, err = bufio.NewReader(conn).Read(buf)
	if err == nil {
		fmt.Printf("message from %s: %s \n", conn.RemoteAddr(), buf)
	} else {
		fmt.Printf("error reading: %v\n", err)
	}
}
