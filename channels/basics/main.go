package main

import "fmt"

func main() {
	// c := make(chan int)
	// c <- 42
	// fmt.Println(<-c)
	// This code doesnt work

	//-----------------------
	//working method 1
	// c := make(chan int)
	// go func() {
	// 	c <- 42
	// }()
	// fmt.Println(<-c)

	//---------------------
	//or another way is method 2
	// c1 := make(chan int, 1)
	// c1 <- 43
	// fmt.Println(<-c1)

	//---------------------
	//Unsuccessful buffers
	// c2 := make(chan int, 1)
	// c2 <- 42
	// c2 <- 43
	// fmt.Println(<-c2)

	//---------------------
	//Successfull buffer
	c3 := make(chan int, 2)
	c3 <- 42
	c3 <- 43
	fmt.Println(<-c3)
	fmt.Println(<-c3)
}
