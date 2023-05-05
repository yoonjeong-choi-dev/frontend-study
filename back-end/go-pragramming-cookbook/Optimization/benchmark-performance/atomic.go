package benchmark_performance

import "sync/atomic"

// AtomicCounter atomic 패지지 이용하여 동시성 제어
type AtomicCounter struct {
	value int64
}

func (c *AtomicCounter) Add(amount int64) {
	atomic.AddInt64(&c.value, amount)
}

func (c *AtomicCounter) Read() int64 {
	var ret int64
	ret = atomic.LoadInt64(&c.value)
	return ret
}
