package middleware

import (
	"net/http"
	"path"
	"strings"
)

// RestrictPrefix 정적 리소스를 요청하는 경우, 특정 prefix 에 대한 필터링
func RestrictPrefix(prefix string, next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			for _, p := range strings.Split(path.Clean(r.URL.Path), "/") {
				if strings.HasPrefix(p, prefix) {
					// prefix 로 시작하는 경로가 있는 경우, 다음 핸들러 적용 X
					// => 보안을 위해 해당 파일이 없다고 응답
					http.Error(w, "Not Found", http.StatusNotFound)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
}
