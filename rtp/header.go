package rtp

// Header represents a RTP packet header.
// 0                   1                   2                   3
// 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |V=2|P|X|  CC   |M|     PT      |       sequence number         |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                           timestamp                           |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |           synchronization source (SSRC) identifier            |
// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+
// |            contributing source (CSRC) identifiers             |
// |                             ....                              |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//
// Header contains first two octets:
// - version (V): 2 bits - Identifies the version of RTP. V=2
// - padding (P): 1 bit - If the padding bit is set, the packet contains
// one or more additional padding octets at the end which are not part
// of the payload.
// - extension (X): 1 bit - If the extension bit is set, the fixed header
// MUST be followed by exactly one header extension.
// CSRC count (CC): 4 bits - The CSRC count contains the number of CSRC
// identifiers that follow the fixed header.
// - marker (M): 1 bit - Intended to allow significant events such as frame
// boundaries to be marked in the packet stream.
// payload type (PT): 7 bits - Identifies the format of the RTP payload
// and determines its interpretation by the application.
//
// SequenceNumber contains next two octets:
// - sequence number: 16 bits - Increments by one for each RTP data packet
// sent, and may be used by the receiver to detect packet loss and to
// restore packet sequence.  The initial value of the sequence number
// SHOULD be random (unpredictable)
//
// Timestamp contains next 4 octets
// - timestamp: 32 bits - Reflects the sampling instant of the first octet in
// the RTP data packet.
//
// SynchronizationSource contains next 4 octets
// - SSRC: 32 bits - Identifies the synchronization source.
//
// - CSRC list: 0 to 15 items, 32 bits each - Identifies the contributing
// sources for the payload contained in this packet.
type Header interface{}

type header struct {
	Head                  uint16   // |V (2 bits)|P (1 bit)|X (1 bit)|CC (4 bits)|M (1 bit)|PT (7 bits)|
	SequenceNumber        uint16   // |sequence number (16 bits)|
	Timestamp             uint32   // |timestamp (32 bits)|
	SynchronizationSource uint32   // |SSRC (32 bits)|
	ContributingSources   []uint32 // |CSRC (array of up to 15 32 bits)|
}

// NewHeader returns a prepared header set with:
// |V |P|X| CC |M|  PT   |
// |10|0|0|0000|0|0000000|
func NewHeader() Header {
	head := uint16(32768) // 0b1000000000000000
	return header{head, 0, 0, 0, []uint32{}}
}

// NextHeader returns a new header with all of the header h's attributes but an
// incremented sequence number
func NextHeader(h Header) Header {
	hdr := h.(header)
	return header{hdr.Head, hdr.SequenceNumber + 1, 0, 0, []uint32{}}
}
