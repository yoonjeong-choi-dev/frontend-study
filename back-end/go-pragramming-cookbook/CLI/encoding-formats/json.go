package encoding_formats

import (
	"bytes"
	"encoding/json"
	"os"
)

type JSONAddress struct {
	City    string `json:"city"`
	Country string `json:"country"`
	IsAlone bool   `json:"is_alone"`
}

type JSONData struct {
	Name    string      `json:"name"`
	Age     int         `json:"age"`
	Address JSONAddress `json:"address"`
}

func (t *JSONData) ToJSON() (*bytes.Buffer, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(b)
	return buffer, nil
}

func (t *JSONData) WriteFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	b, err := json.Marshal(t)
	if err != nil {
		return err
	}

	_, err = file.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func (t *JSONData) Decode(data []byte) error {
	return json.Unmarshal(data, t)
}

func (t *JSONData) ReadFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	decoder := json.NewDecoder(file)
	return decoder.Decode(t)
}
