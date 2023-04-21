package stateful_handler

// Storage 서버와 연결된 스토리지 인터페이스
type Storage interface {
	Get() string
	Put(string)
}

// InMemoryStorage Storage 인터페이스 구현체
type InMemoryStorage struct {
	value string
}

func (m *InMemoryStorage) Get() string {
	return m.value
}

func (m *InMemoryStorage) Put(value string) {
	m.value = value
}
