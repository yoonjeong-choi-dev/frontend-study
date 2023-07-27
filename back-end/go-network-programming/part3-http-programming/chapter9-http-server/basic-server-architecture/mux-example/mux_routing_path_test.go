package mux_example

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// drainAndClose 테스트 용 단발성 요청 Body 스트림을 소비하는 미들웨어
func drainAndClose(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 요청 처리
		next.ServeHTTP(w, r)

		// request Body 스트림 소비
		_, _ = io.Copy(ioutil.Discard, r.Body)
		_ = r.Body.Close()
	})
}

func TestRoutingPathInMux(t *testing.T) {
	mux := http.NewServeMux()

	// 멀티플렉서에 핸들러 등록
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	mux.HandleFunc("/hello", func(w http.ResponseWriter,
		r *http.Request) {
		_, _ = fmt.Fprint(w, "Hello")
	})
	mux.HandleFunc("/hello/sub/", func(w http.ResponseWriter,
		r *http.Request) {
		_, _ = fmt.Fprint(w, "Hello Sub")
	})

	// 미들웨어를 적용한 전체 핸들러
	handler := drainAndClose(mux)

	testCases := []struct {
		path     string
		response string
		code     int
	}{
		{"http://test/", "", http.StatusNoContent},
		{"http://test/hello", "Hello", http.StatusOK},
		{"http://test/hello/sub/", "Hello Sub", http.StatusOK},
		// 서브트리 경로: http://test/hello/sub: redirect to http://test/hello/sub/
		{"http://test/hello/sub",
			"<a href=\"/hello/sub/\">Moved Permanently</a>.\n\n",
			http.StatusMovedPermanently},
		// 서브트리 경로: http://test/hello/sub/* -> http://test/hello/sub/
		{"http://test/hello/sub/you", "Hello Sub", http.StatusOK},
		{"http://test/something/else/entirely", "", http.StatusNoContent},
		{"http://test/hello/you", "", http.StatusNoContent},
	}

	for _, test := range testCases {
		r := httptest.NewRequest(http.MethodGet, test.path, nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)

		res := w.Result()
		got := res.StatusCode
		if got != test.code {
			t.Errorf("[GET %s] expected: %d, got: %d\n", test.path, test.code, got)
		}

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		_ = res.Body.Close()

		body := string(data)
		if body != test.response {
			t.Errorf("[GET %s] expected: %s, got: %s\n", test.path, test.response, body)
		}

		t.Logf("[GET %s] body: %s, code: %d\n", test.path, body, got)
	}
}
