package cps

import "github.com/showbufire/kv"

type entry struct {
	result int
	ready  chan struct{}
}

type getReq struct {
	k  int
	rc chan int
}

type memo struct {
	cached map[int]*entry
	gch    chan *getReq
}

func NewMemo() *memo {
	m := &memo{
		cached: make(map[int]*entry),
		gch:    make(chan *getReq),
	}
	go m.monitor()
	return m
}

func (m *memo) Get(k int) int {
	r := &getReq{
		k:  k,
		rc: make(chan int),
	}
	m.gch <- r
	return <-r.rc
}

func (m *memo) monitor() {
	for {
		select {
		case r := <-m.gch:
			e, ok := m.cached[r.k]
			if !ok {
				e = &entry{ready: make(chan struct{})}
				m.cached[r.k] = e
				go func(e *entry) {
					e.result = kv.ExpensiveSquare(r.k)
					close(e.ready)
				}(e)
			}
			go func() {
				<-e.ready
				r.rc <- e.result
			}()
		}
	}
}
