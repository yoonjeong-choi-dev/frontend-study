package encoding

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestInvalidPayloadSize(t *testing.T) {
	buf := new(bytes.Buffer)
	err := buf.WriteByte(BinaryType)
	if err != nil {
		t.Fatalf("Error for write Type Header: %s\n", err.Error())
	}

	// Set length to write 1 GB message
	err = binary.Write(buf, binary.BigEndian, uint32(1<<30))
	if err != nil {
		t.Fatalf("Error for write Length Header: %s\n", err.Error())
	}

	// Test
	var b Binary
	_, err = b.ReadFrom(buf)
	if err != ErrMaxPayloadSize {
		t.Fatalf("expected ErrMaxPayloadSize, got %v\n", err)
	}
}
