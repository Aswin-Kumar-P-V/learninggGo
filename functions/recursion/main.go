package main

import "fmt"

func main() {
	x := factorial(5)
	fmt.Println(x)
}
func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}
