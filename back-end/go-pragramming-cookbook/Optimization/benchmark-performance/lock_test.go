package benchmark_performance

import (
	"sync"
	"testing"
)

func BenchmarkCounterAdd(b *testing.B) {
	c := Counter{value: 0, mutex: &sync.RWMutex{}}
	for n := 0; n < b.N; n++ {
		c.Add(1)
	}
}

func BenchmarkCounterRead(b *testing.B) {
	c := Counter{value: 0, mutex: &sync.RWMutex{}}
	for n := 0; n < b.N; n++ {
		c.Read()
	}
}

func BenchmarkCounterAddRead(b *testing.B) {
	c := Counter{value: 0, mutex: &sync.RWMutex{}}
	for n := 0; n < b.N; n++ {
		c.Add(1)
		c.Read()
	}
}
