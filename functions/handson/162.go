package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello world")
	}()

	x := func() {
		fmt.Println("Call from x")
	}
	x()
	y := foo1()
	y()
}
func foo1() func() {
	return func() {
		fmt.Println("Hello from func inside func")
	}
}
