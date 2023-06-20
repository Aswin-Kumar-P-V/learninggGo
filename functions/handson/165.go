package main

import "fmt"

func square1(n int) int {
	return n * n
}
func printSquare(s func(int) int, a int) {
	x := s(a)
	fmt.Printf("square of %v is %v", a, x)
}
func main() {
	printSquare(square1, 10)
}
