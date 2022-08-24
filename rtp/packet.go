package rtp

// Packet represents an RTP packet
type Packet struct {
	Header
	Payload []byte
}
