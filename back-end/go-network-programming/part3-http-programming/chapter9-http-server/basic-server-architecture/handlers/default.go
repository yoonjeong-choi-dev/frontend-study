package handlers

import (
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
)

// greetTemplate template 패키지를 이용하여 응답 데이터의 문자열 이스케이핑
var greetTemplate = template.Must(template.New("hello").Parse("Hello, {{.}}!"))

func DefaultHandler() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func(r io.ReadCloser) {
				_, _ = io.Copy(ioutil.Discard, r)
				_ = r.Close()
			}(r.Body)

			var resBuffer []byte

			switch r.Method {
			case http.MethodGet:
				resBuffer = []byte("anonymous")
			case http.MethodPost:
				var err error

				// 요청 데이터가 그대로 응답으로 쓰이는 상황
				// => 요청 데이터(문자열)에 대한 이스케이핑을 하여 XXS 방지해야 함
				resBuffer, err = ioutil.ReadAll(r.Body)
				if err != nil {
					http.Error(w, "error for reading request body", http.StatusInternalServerError)
					return
				}
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			_ = greetTemplate.Execute(w, string(resBuffer))
		})
}
