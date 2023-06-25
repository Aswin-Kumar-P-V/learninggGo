package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 42
	c <- 43
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("----------------")
	fmt.Printf("%T\n", c)
	fmt.Println("----------******---------")
	//This code is ok as we can send and recieve from our channel c

	//lets see a one directional channel

	c1 := make(chan<- int, 1)
	c1 <- 42
	// fmt.Println(<-c1)
	//This print statement brings error as it is a recieve on channel

	fmt.Println("----------*******----------")

	//Another example

	c2 := make(<-chan int, 1)
	// c2 <- 42 makes error as it is not a reciever channel
	fmt.Println(<-c2)
}
