package lock

import (
	"sync"

	"github.com/showbufire/kv"
)

type entry struct {
	result int
	ready  chan int
}

type memo struct {
	cached map[int]*entry
	mtx    *sync.Mutex
}

func NewMemo() *memo {
	return &memo{
		cached: make(map[int]*entry),
		mtx:    &sync.Mutex{},
	}
}

func (m *memo) Get(k int) int {
	m.mtx.Lock()
	v, ok := m.cached[k]
	if !ok {
		ready := make(chan int)
		v = &entry{result: -1, ready: ready}
		m.cached[k] = v
		m.mtx.Unlock()
		v.result = kv.ExpensiveSquare(k)
		close(v.ready)
	} else {
		m.mtx.Unlock()
		<-v.ready
	}
	return v.result
}
