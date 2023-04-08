package null_encoding

import (
	"encoding/json"
	"fmt"
)

const (
	jsonBlob     = `{"name": "Yoonjeong"}`
	fullJsonBlob = `{"name": "YJ", "age":0}`
)

type BasicData struct {
	// omitempty 옵션은 0인 값을 인코딩하지 못함
	Age  int    `json:"age,omitempty"`
	Name string `json:"name"`
}

func BaseEncoding() error {
	d := BasicData{}
	if err := json.Unmarshal([]byte(jsonBlob), &d); err != nil {
		return err
	}

	fmt.Printf("Basic json data with no age: %s\n", jsonBlob)
	fmt.Printf("Unmarshal Result: %+v\n", d)

	val, err := json.Marshal(d)
	if err != nil {
		return err
	}
	fmt.Printf("Marshal Result with no age: %s\n", string(val))

	if err := json.Unmarshal([]byte(fullJsonBlob), &d); err != nil {
		return err
	}
	fmt.Printf("Basic json data with age=0: %s\n", fullJsonBlob)
	fmt.Printf("Unmarshal Result: %+v\n", d)

	val, err = json.Marshal(d)
	if err != nil {
		return err
	}
	fmt.Printf("Marshal Result with age=0: %s\n", string(val))

	return nil
}
