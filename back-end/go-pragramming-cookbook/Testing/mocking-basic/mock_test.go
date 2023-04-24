package mocking

type MockDoStuffer struct {
	MockDoStuff func(input string) error
}

// DoStuff DoStuffer 인터페이스 구현
// : 이때 모킹 함수인 MockDoStuff 함수가 대리자 역할
func (m *MockDoStuffer) DoStuff(input string) error {
	if m.MockDoStuff != nil {
		return m.MockDoStuff(input)
	}

	return nil
}
