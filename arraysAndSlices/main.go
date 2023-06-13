package main

import "fmt"

func main() {
	a := [2]int{0, 1}
	fmt.Printf("Array a is %v \n", a)
	b := [...]string{"Hello", "world"}
	fmt.Printf("Array b is %v \n", b)

	var c [2]int
	c[0] = 1
	c[1] = 2
	fmt.Printf("Array c is %v\n", c)
}
