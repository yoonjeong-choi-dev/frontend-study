package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRestrictPrefix(t *testing.T) {
	// static 경로는 ../files/ 디렉터리로 매핑하여 정적 파일 제공
	handler := http.StripPrefix("/static/",
		RestrictPrefix(".", http.FileServer(http.Dir("../files/"))),
	)

	testCases := []struct {
		path string
		code int
	}{
		{"http://test/static/sage.svg", http.StatusOK},
		{"http://test/static/.secret", http.StatusNotFound},
		{"http://test/static/.dir/secret", http.StatusNotFound},
	}

	for i, test := range testCases {
		r := httptest.NewRequest(http.MethodGet, test.path, nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)

		got := w.Result().StatusCode
		if test.code != got {
			t.Errorf("[%d] Expected: %d, Got: %d\n", i, test.code, got)
		}
		t.Logf("%s -> %d\n", test.path, got)
	}
}
