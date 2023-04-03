package encoding_formats

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"os"
)

type TOMLAddress struct {
	City    string `toml:"city"`
	Country string `toml:"country"`
	IsAlone bool   `toml:"is_alone"`
}

type TOMLData struct {
	Name    string      `toml:"name"`
	Age     int         `toml:"age"`
	Address TOMLAddress `toml:"address"`
}

func (t *TOMLData) ToTOML() (*bytes.Buffer, error) {
	b := &bytes.Buffer{}

	encoder := toml.NewEncoder(b)
	if err := encoder.Encode(t); err != nil {
		return nil, err
	}

	return b, nil
}

func (t *TOMLData) WriteFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	encoder := toml.NewEncoder(file)
	if err := encoder.Encode(t); err != nil {
		return err
	}
	return nil
}

func (t *TOMLData) Decode(data []byte) (toml.MetaData, error) {
	return toml.Decode(string(data), t)
}

func (t *TOMLData) ReadFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	_, err = toml.DecodeFile(fileName, t)
	return err
}
