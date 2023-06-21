package main

import "fmt"

func main() {
	a := 42
	IncrementNormal(a)
	fmt.Println(a)
	IncrementPointer(&a)
	fmt.Println(a)
	is := []int{1, 2, 3, 4}
	fmt.Println(is)
	ChSlice(is)
	fmt.Println(is)
}

func IncrementNormal(n int) {
	n++
}

func IncrementPointer(n *int) {
	*n++
}

func ChSlice(ii []int) {
	ii[0] = 99
}
