package main

import "fmt"

func foo() int {
	a := 42
	return a
}

func bar() (string, int) {
	return "Hello from boo", 23
}

func main() {
	a := foo()
	s, b := bar()
	fmt.Println(a)
	fmt.Println(s, b)
}
