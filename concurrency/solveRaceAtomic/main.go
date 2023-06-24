package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUS:", runtime.NumCPU())
	fmt.Println("Go Routines:", runtime.NumGoroutine())
	var counter int64
	const gs = 50
	var wG sync.WaitGroup
	wG.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter: ", atomic.LoadInt64(&counter))
			runtime.Gosched()

			wG.Done()
		}()
	}
	wG.Wait()
	fmt.Println("Go Routines:", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
}
