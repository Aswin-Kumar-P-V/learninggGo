package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	for i := 0; i < 10; i++ {

		go func(i int) {
			c <- i
		}(i)
	}

	for j := 0; j < 10; j++ {
		fmt.Println(<-c)
	}
	close(c)
}
