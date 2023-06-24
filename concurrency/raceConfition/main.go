package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUS:", runtime.NumCPU())
	fmt.Println("Go Routines:", runtime.NumGoroutine())
	counter := 0
	const gs = 50
	var wG sync.WaitGroup
	wG.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			inc := counter
			runtime.Gosched()
			inc++
			counter = inc
			wG.Done()
		}()
	}
	wG.Wait()
	fmt.Println("Go Routines:", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
}
