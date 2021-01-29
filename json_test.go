package goson

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestJson(t *testing.T) {
	var v int64 = 0
	wg := &sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			atomic.AddInt64(&v, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	t.Log(v)

	fmt.Println()
}