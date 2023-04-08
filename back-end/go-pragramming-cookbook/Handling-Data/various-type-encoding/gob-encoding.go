package various_type_encoding

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Position struct {
	X      int
	Y      int
	Object string
}

func GobExample() error {
	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)

	p := Position{
		X:      10,
		Y:      10,
		Object: "Toy",
	}

	if err := encoder.Encode(&p); err != nil {
		return err
	}

	fmt.Printf("Gob Encoded Data Length: %d\n", len(buffer.Bytes()))

	decoded := Position{}

	// 버퍼에 저장된 인코딩된 바이너리를 디코더에 등록
	decoder := gob.NewDecoder(&buffer)
	if err := decoder.Decode(&decoded); err != nil {
		return err
	}
	fmt.Printf("Gob Decoded Data: %+v\n", decoded)

	return nil
}
