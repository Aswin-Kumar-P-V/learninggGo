package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	count := 0
	var wg sync.WaitGroup
	const gs = 50
	var mu sync.Mutex
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			inc := count
			runtime.Gosched()
			inc++
			count = inc
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
