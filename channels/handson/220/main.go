package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 42
	}()

	v, ok := <-c
	if ok {
		fmt.Println(v, ok)
	}

	close(c)

	v, ok = <-c

	fmt.Println(v, ok)
}
