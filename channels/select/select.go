package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(eve, odd, quit)

	recieve(eve, odd, quit)
	fmt.Println("Exit")
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func recieve(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Value recieved from eve channel:", v)
		case v := <-o:
			fmt.Println("Value recieved from odd channel:", v)
		case v := <-q:
			fmt.Println("Value recieved from quit channel:", v)
			return
		}
	}
}
