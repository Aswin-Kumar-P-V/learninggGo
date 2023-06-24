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
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			inc := count
			runtime.Gosched()
			inc++
			count = inc
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
