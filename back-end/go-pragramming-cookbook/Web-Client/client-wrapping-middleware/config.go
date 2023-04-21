package wrapping

import (
	"log"
	"net/http"
	"os"
)

func Setup() *http.Client {
	client := http.Client{}

	// 미들웨어 등록
	transport := Decorate(
		&http.Transport{},
		Logger(log.New(os.Stdout, "[Logger]", 0)),
		BasicAuth("yoonjeong-choi", "yj-password"),
	)

	client.Transport = transport

	return &client
}
