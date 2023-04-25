package atomic

import (
	"sync"
	"sync/atomic"
)

// Ordinal 한 번만 초기화할 수 있는 전역 데이터
type Ordinal struct {
	ordinal uint64
	once    *sync.Once
}

func (o *Ordinal) Init(val uint64) {
	o.once.Do(func() {
		atomic.StoreUint64(&o.ordinal, val)
	})
}

func (o *Ordinal) Get() uint64 {
	// atomic 패키지를 이용하여, 원자적 처리
	// : mutex 처리와 동일
	return atomic.LoadUint64(&o.ordinal)
}

func (o *Ordinal) Inc() {
	// atomic 패키지를 이용하여, 원자적 처리
	atomic.AddUint64(&o.ordinal, 1)
}

func NewOrdinal() *Ordinal {
	return &Ordinal{
		once: &sync.Once{},
	}
}
