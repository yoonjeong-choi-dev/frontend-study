package decoder

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"type-length-value-encoding/encoding"
)

func Decoder(r io.Reader) (encoding.Payload, error) {
	var typ uint8

	err := binary.Read(r, binary.BigEndian, &typ)
	if err != nil {
		return nil, err
	}

	var payload encoding.Payload
	switch typ {
	case encoding.BinaryType:
		payload = new(encoding.Binary)
	case encoding.StringType:
		payload = new(encoding.String)
	default:
		return nil, errors.New("unknown message type")
	}

	// 이미 타입 추론을 위해 1바이트를 읽었기 때문에, ReadFrom 메서드에 타입 정보를 추가하여 전달해야 함
	// => ReadFrom 메서드에서 타입 관련 정보를 읽지 않게 하는 방법으로 리팩터링 가능
	_, err = payload.ReadFrom(
		io.MultiReader(bytes.NewReader([]byte{typ}), r))
	if err != nil {
		return nil, err
	}
	return payload, nil
}
