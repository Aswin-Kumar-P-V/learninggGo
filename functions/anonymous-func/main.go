package main

import "fmt"

func bar() func() int {
	return func() int {
		return 42
	}
}

func main() {
	// func() {
	// 	fmt.Println("HELLO WORLD")
	// }()
	// x := func(s string) {
	// 	fmt.Println("from x : %v", s)
	// }
	// x("Aswin says hi")

	z := bar()
	fmt.Println(z())
}
