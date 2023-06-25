package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(eve, odd, quit)

	recieve(eve, odd, quit)
}

func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- false
	close(q)
}

func recieve(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("From eve channel:", v)
		case v := <-o:
			fmt.Println("From odd channel:", v)
		case v, ok := <-q:
			if ok {
				fmt.Println("From quit channel:", v)
				return
			}

		}
	}
}
