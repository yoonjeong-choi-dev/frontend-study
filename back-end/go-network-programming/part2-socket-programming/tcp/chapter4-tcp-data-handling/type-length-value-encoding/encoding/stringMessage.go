package encoding

import (
	"encoding/binary"
	"errors"
	"io"
)

// String implementation of Payload for String type message
type String string

func (m *String) Bytes() []byte  { return []byte(*m) }
func (m *String) String() string { return string(*m) }

func (m *String) WriteTo(w io.Writer) (int64, error) {
	// Write Message Type Header
	err := binary.Write(w, binary.BigEndian, StringType)
	if err != nil {
		return 0, err
	}

	// Write Message Length Header
	var size int64 = HeaderForTypeBytes
	err = binary.Write(w, binary.BigEndian, uint32(len(*m)))
	if err != nil {
		return size, err
	}

	// Write Message Body
	size += HeaderForLengthBytes

	bodySize, err := w.Write(m.Bytes())
	if err != nil {
		return size, err
	}
	return size + int64(bodySize), err
}

func (m *String) ReadFrom(r io.Reader) (int64, error) {
	// Read Message Type Header
	var typ uint8
	err := binary.Read(r, binary.BigEndian, &typ)
	if err != nil {
		return 0, err
	}

	var size int64 = HeaderForTypeBytes

	// Validate Message Type
	if typ != StringType {
		return size, errors.New("invalid message type: must be StringType")
	}

	// Read Message Length Header
	var msgSize uint32
	err = binary.Read(r, binary.BigEndian, &msgSize)
	if err != nil {
		return size, err
	}

	// Validate Message Length
	size += HeaderForLengthBytes
	if msgSize > MaxPayloadSize {
		return size, ErrMaxPayloadSize
	}

	// Read Message Body
	// 헤더의 길이 정보를 이용하여 버퍼 사이즈 설정
	buf := make([]byte, msgSize)
	bodySize, err := r.Read(buf)
	if err != nil {
		return size, err
	}
	*m = String(buf)
	return size + int64(bodySize), nil
}
