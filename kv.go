package kv

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Memo interface {
	Get(k int) int
}

func ExpensiveSquare(k int) int {
	time.Sleep(time.Second)
	return k * k
}

func shuffle(a []int) {
	for i := 0; i < len(a); i += 1 {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func Run(m Memo, nthreads, nsize int) {
	rand.Seed(time.Now().Unix())
	wg := &sync.WaitGroup{}
	for i := 0; i < nthreads; i += 1 {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			ks := make([]int, nsize, nsize)
			for j := 0; j < nsize; j += 1 {
				ks[j] = j
			}
			shuffle(ks)
			for _, k := range ks {
				v := m.Get(k)
				fmt.Printf("thread %d query %d got %d\n", idx, k, v)
			}
		}(i)
	}
	wg.Wait()
}
