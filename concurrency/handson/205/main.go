package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count int64
	var wg sync.WaitGroup
	const gs = 50
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			atomic.LoadInt64(&count)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
