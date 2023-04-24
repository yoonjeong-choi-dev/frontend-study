package auto_mock_gen

// GetSetter 테스트에 사용할 인터페이스
type GetSetter interface {
	Set(key, val string) error
	Get(key string) (string, error)
}
