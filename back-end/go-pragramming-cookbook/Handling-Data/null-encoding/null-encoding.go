package null_encoding

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

// nullInt64 custom type for encoding/json
// => implement marshal and unmarshal function to use json package
type nullInt64 sql.NullInt64

func (v *nullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	}
	return json.Marshal(nil)
}

func (v *nullInt64) UnmarshalJSON(b []byte) error {
	v.Valid = false
	if b != nil {
		v.Valid = true
		return json.Unmarshal(b, &v.Int64)
	}
	return nil
}

type NullIntData struct {
	Age  *nullInt64 `json:"age,omitempty"`
	Name string     `json:"name"`
}

func CustomNullTypeExample() error {
	d := NullIntData{}
	if err := json.Unmarshal([]byte(jsonBlob), &d); err != nil {
		return err
	}

	fmt.Printf("Custom Null json data with no age: %s\n", jsonBlob)
	fmt.Printf("Custom Null Unmarshal Result: %+v\n", d)

	val, err := json.Marshal(d)
	if err != nil {
		return err
	}
	fmt.Printf("Custom Null Result with no age: %s\n", string(val))

	if err := json.Unmarshal([]byte(fullJsonBlob), &d); err != nil {
		return err
	}
	fmt.Printf("Custom Null json data with age=0: %s\n", fullJsonBlob)
	fmt.Printf("Custom Null Unmarshal Result: %+v\n", d)

	val, err = json.Marshal(d)
	if err != nil {
		return err
	}
	fmt.Printf("Custom Null Marshal Result with age=0: %s\n", string(val))
	return nil
}
