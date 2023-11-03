package atomicpkg

import (
	"sync"
	"sync/atomic"
)

type Ordinal struct {
	ordinal uint64
	once    *sync.Once
}

func NewOrdinal() *Ordinal {
	return &Ordinal{once: &sync.Once{}}
}

func (o *Ordinal) Init(val uint64) {
	o.once.Do(func() {
		atomic.StoreUint64(&o.ordinal, val)
	})
}

// 현재 ordinal 값 반환
func (o *Ordinal) GetOrdinal() uint64 {
	return atomic.LoadUint64(&o.ordinal)
}

// 현재 ordinal 값 증가
func (o *Ordinal) Increment() {
	atomic.AddUint64(&o.ordinal, 1)
}
