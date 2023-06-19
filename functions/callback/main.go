package main

import "fmt"

func main() {
	x := doMath(10, 20, add)
	y := doMath(20, 5, sub)
	fmt.Println("sum:", x)
	fmt.Println("diff:", y)
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}
