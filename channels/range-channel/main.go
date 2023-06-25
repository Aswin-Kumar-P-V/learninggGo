package main

import (
	"fmt"
	"runtime"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	fmt.Println("goro", runtime.NumGoroutine())
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("Exit")
}
