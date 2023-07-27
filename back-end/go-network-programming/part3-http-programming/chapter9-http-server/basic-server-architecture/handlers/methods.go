package handlers

import (
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

type Methods map[string]http.Handler

// ServerHTTP implementation of http.Handler interface
// => Methods 타입 자체가 핸들러로 사용 가능
func (m Methods) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func(r io.ReadCloser) {
		_, _ = io.Copy(ioutil.Discard, r)
		_ = r.Close()
	}(r.Body)

	if handler, ok := m[r.Method]; ok {
		if handler == nil {
			http.Error(w, "cannot parsing possible method", http.StatusInternalServerError)
		} else {
			// 요청 메서드를 확인하여, 대응하는 핸들러로 요청 라우팅
			// => Methods 타입은 멀티플렉서(라우터) 역할
			handler.ServeHTTP(w, r)
		}
		return
	}

	// 요청 메서드는 지원하지 않음 => 지원하는 메서드들을 함께 전송
	w.Header().Add("Allow", m.allowedMethods())
	if r.Method != http.MethodOptions {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (m Methods) allowedMethods() string {
	ret := make([]string, 0, len(m))

	for method := range m {
		ret = append(ret, method)
	}

	sort.Strings(ret)
	return strings.Join(ret, ", ")
}

// DefaultMethodsHandler DefaultHandler 핸들러를 개선
// => 지원하는 메서드들 전송 및 OPTIONS 메서드 지원
func DefaultMethodsHandler() http.Handler {
	// Methods.ServeHTTP 에서 응답 스트림 닫기 처리를 하므로,
	// 여기서는 요청 자체에 대한 로직에 집중 가능
	return Methods{
		http.MethodGet: http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write([]byte("Hello, anonymous!"))
			},
		),
		http.MethodPost: http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				resBuffer, err := ioutil.ReadAll(r.Body)
				if err != nil {
					http.Error(w, "error for reading request body",
						http.StatusInternalServerError)
					return
				}

				_, _ = fmt.Fprintf(w, "Hello, %s!",
					html.EscapeString(string(resBuffer)))
			},
		),
	}
}
