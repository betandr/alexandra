package client

import (
	"fmt"
	"net"

	"github.com/betandr/alexandra/rtp"
)

// NewConnection starts a UDP client on port 8067
func NewConnection() (net.Conn, error) {
	conn, err := net.Dial("udp", "127.0.0.1:8067")
	defer conn.Close()
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// Send
func Send(conn net.Conn, pkt rtp.Packet) {
	fmt.Fprintf(conn, pkt.Payload)
}
