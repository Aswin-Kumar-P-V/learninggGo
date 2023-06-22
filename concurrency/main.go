package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS:\t", runtime.GOOS)
	fmt.Println("ARCH:\t", runtime.GOARCH)
	fmt.Println("NO CPU's:\t", runtime.NumCPU())
	fmt.Println("Go Routines:\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo: ", i)
	}
	wg.Done()
}

func bar() {
	j := 0
	for j < 10 {
		fmt.Println("Bar: ", j)
		j++
	}
}
