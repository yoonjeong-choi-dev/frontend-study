package encoding_formats

import (
	"bytes"
	"github.com/go-yaml/yaml"
	"os"
)

type YAMLAddress struct {
	City    string `json:"city"`
	Country string `json:"country"`
	IsAlone bool   `json:"is-alone"`
}

type YAMLData struct {
	Name    string      `json:"name"`
	Age     int         `json:"age"`
	Address YAMLAddress `json:"address"`
}

func (t *YAMLData) ToYAML() (*bytes.Buffer, error) {
	b, err := yaml.Marshal(t)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(b)
	return buffer, nil
}

func (t *YAMLData) WriteFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	b, err := yaml.Marshal(t)
	if err != nil {
		return err
	}

	_, err = file.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func (t *YAMLData) Decode(data []byte) error {
	return yaml.Unmarshal(data, t)
}

func (t *YAMLData) ReadFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	decoder := yaml.NewDecoder(file)
	return decoder.Decode(t)
}
