package main

import (
	"fmt"
	"sync"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	fan := make(chan int)

	go send(eve, odd)

	go recieve(eve, odd, fan)

	for v := range fan {
		fmt.Println(v)
	}

}

func send(e, o chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

func recieve(e, o <-chan int, fan chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range e {
			fan <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range o {
			fan <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fan)
}
