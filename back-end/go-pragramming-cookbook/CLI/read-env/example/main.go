package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	read_env "read-env"
)

type Config struct {
	Version string `json:"version" required:"true"`
	IsSafe  bool   `json:"is_safe" required:"true"`
	Secret  string `json:"secret"`
}

func main() {
	var err error

	// 테스트를 위한 임시 json 파일 생성
	tempFile, err := ioutil.TempFile("", "tmp")
	if err != nil {
		log.Fatalf("Error for creating temp file: %s\n", err.Error())
	}
	defer func() {
		_ = tempFile.Close()
		os.Remove(tempFile.Name())
	}()

	// 임시 설정 파일에 저장할 데이터
	secrets := `{
        "secret": "YJ Secret Key!"
    }`

	if _, err = tempFile.Write(bytes.NewBufferString(secrets).Bytes()); err != nil {
		log.Fatalf("Error for writing config file: %s\n", err.Error())
	}

	// Set Environment Variables
	const envPrefix string = "EXAMPLE_ENV"
	versionEnv := fmt.Sprintf("%s_VERSION", envPrefix)
	isSaveEnv := fmt.Sprintf("%s_ISSAFE", envPrefix)
	if err = os.Setenv(versionEnv, "1.7.3"); err != nil {
		panic(err)
	}
	if err = os.Setenv(isSaveEnv, "true"); err != nil {
		panic(err)
	}

	// Load Environment Variables and Config Files
	c := Config{}
	if err = read_env.LoadConfig(tempFile.Name(), envPrefix, &c); err != nil {
		panic(err)
	}

	// Compare config with env
	fmt.Printf("version: %s vs %s\n", c.Version, os.Getenv(versionEnv))
	fmt.Printf("isSafe: %v vs %s\n", c.IsSafe, os.Getenv(isSaveEnv))
	fmt.Printf("secret from file: %s\n", c.Secret)
	fmt.Printf("Loaded Config: %#v\n", c)
}
