package encoding

import (
	"encoding/binary"
	"errors"
	"io"
)

// Binary implementation of Payload for Binary type message
type Binary []byte

func (m *Binary) Bytes() []byte  { return *m }
func (m *Binary) String() string { return string(*m) }

func (m *Binary) WriteTo(w io.Writer) (int64, error) {
	// Write Message Type Header
	err := binary.Write(w, binary.BigEndian, BinaryType)
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

	bodySize, err := w.Write(*m)
	if err != nil {
		return size, err
	}
	return size + int64(bodySize), err
}

func (m *Binary) ReadFrom(r io.Reader) (int64, error) {
	// Read Message Type Header
	var typ uint8
	err := binary.Read(r, binary.BigEndian, &typ)
	if err != nil {
		return 0, err
	}

	var size int64 = HeaderForTypeBytes

	// Validate Message Type
	if typ != BinaryType {
		return size, errors.New("invalid message type: must be BinaryType")
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
	*m = make([]byte, msgSize)
	bodySize, err := r.Read(*m)
	return size + int64(bodySize), err
}
