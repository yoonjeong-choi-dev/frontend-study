package various_type_encoding

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

var dataStr = "Data for Encoding with base64"

func Base64Example() error {
	fmt.Printf("Test String: %s\n", dataStr)
	data := []byte(dataStr)

	encoded := base64.URLEncoding.EncodeToString(data)
	fmt.Printf("With URL Encoding: %s\n", encoded)

	decoded, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return err
	}

	fmt.Printf("With URL Decoding: %s\n", decoded)

	return nil
}

func Base64ExampleWithEncoder() error {
	fmt.Printf("Test String: %s\n", dataStr)
	data := []byte(dataStr)

	buffer := bytes.Buffer{}
	encoder := base64.NewEncoder(base64.StdEncoding, &buffer)

	if _, err := encoder.Write(data); err != nil {
		return nil
	}

	if err := encoder.Close(); err != nil {
		return err
	}

	fmt.Printf("With Standard Encoding: %s\n", buffer.String())

	decoder := base64.NewDecoder(base64.StdEncoding, &buffer)
	decoded, err := ioutil.ReadAll(decoder)
	if err != nil {
		return nil
	}

	fmt.Printf("With Standard Decoding: %s\n", string(decoded))

	return nil
}
