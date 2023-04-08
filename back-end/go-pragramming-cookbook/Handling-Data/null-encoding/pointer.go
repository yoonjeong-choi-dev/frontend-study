package null_encoding

import (
	"encoding/json"
	"fmt"
)

type PointerData struct {
	// omitempty 옵션은 0인 값을 nil
	Age  *int   `json:"age,omitempty"`
	Name string `json:"name"`
}

func PointerExample() error {
	d := PointerData{}
	if err := json.Unmarshal([]byte(jsonBlob), &d); err != nil {
		return err
	}

	fmt.Printf("Pointer json data with no age: %s\n", jsonBlob)
	fmt.Printf("Pointer Unmarshal Result: %+v\n", d)

	val, err := json.Marshal(d)
	if err != nil {
		return err
	}
	fmt.Printf("Pointer Marshal Result with no age: %s\n", string(val))

	if err := json.Unmarshal([]byte(fullJsonBlob), &d); err != nil {
		return err
	}
	fmt.Printf("Pointer json data with age=0: %s\n", fullJsonBlob)
	fmt.Printf("Pointer Unmarshal Result: %+v\n", d)

	val, err = json.Marshal(d)
	if err != nil {
		return err
	}
	fmt.Printf("Pointer Marshal Result with age=0: %s\n", string(val))
	return nil
}
