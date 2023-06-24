package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("Go routines: ", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("Hello from routine 1")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hello from routine 2")
		wg.Done()
	}()

	fmt.Println("Go routines: ", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Go routines: ", runtime.NumGoroutine())
}
