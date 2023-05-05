package encoding

import (
	"errors"
	"fmt"
	"io"
)

// Define Message Type and Payload Size
const (
	BinaryType uint8 = iota + 1
	StringType
	// MaxPayloadSize for security problem
	MaxPayloadSize uint32 = 10 << 20
)

var ErrMaxPayloadSize = errors.New("maximum payload size exceeded")

// Define TLV Encoding Header
// 데이터 유형:1바이트, 데이터 길이 정보: 4바이트
const (
	HeaderForTypeBytes   = 1
	HeaderForLengthBytes = 4
)

// Payload interface for each message type
type Payload interface {
	io.ReaderFrom  // read from io.Reader
	io.WriterTo    // writer to io.Writer
	Bytes() []byte // convert to byte array
	fmt.Stringer   // convert to string
}
