package read_env

import (
	"encoding/json"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"os"
)

// LoadConfig Load environment values to config
// 1. read json file from 'path' and load
// 2. read os EVN via envPrefix
func LoadConfig(path, envPrefix string, config interface{}) error {
	if path != "" {
		err := LoadFile(path, config)
		if err != nil {
			return errors.Wrap(err, "error loading config from file")
		}
	}
	err := envconfig.Process(envPrefix, config)
	return errors.Wrap(err, "error loading config from env")
}

// LoadFile convert json file to config(interface{})
func LoadFile(path string, config interface{}) (err error) {
	configFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		err = configFile.Close()
		if err != nil {
			return
		}
	}()

	decoder := json.NewDecoder(configFile)
	if err = decoder.Decode(config); err != nil {
		err = errors.Wrap(err, "failed to read config file")
		return err
	}

	return nil
}
