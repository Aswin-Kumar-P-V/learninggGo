package main

import "fmt"

func main() {
	y := increment()
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())

}

func increment() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
