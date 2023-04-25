package atomic

import (
	"errors"
	"sync"
)

type SafeMap struct {
	data  map[string]string
	mutex *sync.RWMutex
}

func (m *SafeMap) Set(key, value string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.data[key] = value
}

func (m *SafeMap) Get(key string) (string, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if val, ok := m.data[key]; ok {
		return val, nil
	}

	return "", errors.New("key not found")
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data:  make(map[string]string),
		mutex: &sync.RWMutex{},
	}
}
