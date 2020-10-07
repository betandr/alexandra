package udp

import (
	"bufio"
	"fmt"
	"net"

	"github.com/betandr/alexandra/rtp"
)

// Client is a UDP sender
type Client struct {
	conn      net.Conn
	connected bool
}

// Connect to UDP addr
func (c *Client) Connect(addr string) error {
	var err error
	c.conn, err = net.Dial("udp", addr)
	if err != nil {
		return err
	}
	c.connected = true
	return nil
}

// Disconnect client from UDP connection
func (c *Client) Disconnect() {
	c.conn.Close()
}

// Send a string message on UDP port 8067
func (c *Client) Send(pkt rtp.Packet) (string, error) {
	p := make([]byte, 2048)
	fmt.Fprintf(c.conn, pkt.Payload)
	var err error
	_, err = bufio.NewReader(c.conn).Read(p)
	if err != nil {
		return "", err
	}

	return string(p), nil
}
