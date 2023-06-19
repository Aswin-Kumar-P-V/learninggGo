package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)
	y := func() int {
		return 43
	}()
	fmt.Println(y)
	z := bar()
	fmt.Println(z())
}
func foo() int {
	return 42
}
func bar() func() int {
	return func() int {
		return 44
	}
}
