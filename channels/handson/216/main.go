package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println("Buffered channel c:", <-c)

	//or

	c1 := make(chan int)

	go func() {
		c1 <- 43
	}()
	fmt.Println("Func c1 :", <-c1)
}
