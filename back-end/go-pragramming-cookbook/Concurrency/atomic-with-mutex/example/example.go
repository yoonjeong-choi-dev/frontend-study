package main

import (
	"atomic"
	"fmt"
	"sync"
)

func main() {
	o := atomic.NewOrdinal()
	m := atomic.NewSafeMap()

	for i := 0; i < 10; i++ {
		// Init 여러번 호출해도 처음 호출만 처리
		o.Init(uint64(123 + i))
	}
	fmt.Printf("Initial Ordinal: %d\n", o.Get())

	// 10 concurrency
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()

			m.Set(fmt.Sprintf("%d", idx), "success")
			o.Inc()
		}(i)
	}
	wg.Wait()

	// Result Check
	for i := 0; i < 10; i++ {
		val, err := m.Get(fmt.Sprintf("%d", i))
		if err != nil || val != "success" {
			panic(err)
		}
	}

	fmt.Printf("Final Ordinal: %d\n", o.Get())
	fmt.Println("Success to process in SafeMap")
}
