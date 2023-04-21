package wrapping

import "net/http"

// TransportFunc RoundTripper 인터페이스를 래핑하는 함수 타입
// => RoundTripper.RoundTrip 함수의 시그니처
type TransportFunc func(r *http.Request) (*http.Response, error)

func (tf TransportFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return tf(r)
}

// Decorator 미들웨어들을 연쇄적으로 등록하기 위한 유틸 타입
// 1) 각 미들웨어는 http.RoundTripper 인터페이스 만족
// => 이때 http.RoundTripper 인터페이스를 래핑한 TransportFunc 타입 사용
// 2) 각 미들웨어는 http.RoundTripper 객체를 반환 => 미들웨어 체이닝에 사용
type Decorator func(tripper http.RoundTripper) http.RoundTripper

// Decorate 미들웨어 체이닝을 위한 유틸함수
func Decorate(t http.RoundTripper, middlewares ...Decorator) http.RoundTripper {
	decorated := t
	for _, middleware := range middlewares {
		decorated = middleware(decorated)
	}
	return decorated
}
