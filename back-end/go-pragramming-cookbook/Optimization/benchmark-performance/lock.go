package benchmark_performance

import "sync"

// Counter 뮤텍스틀 이용하여 동시성 제어
type Counter struct {
	value int64
	mutex *sync.RWMutex
}

func (c *Counter) Add(amount int64) {
	c.mutex.Lock()
	c.value += amount
	c.mutex.Unlock()
}

func (c *Counter) Read() int64 {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.value
}
