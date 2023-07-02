package main

import "fmt"

func main() {
	fmt.Println("2 + 4 = ", sum(2, 4))
	fmt.Println("3 + 5 + 6 = ", sum(3, 5, 6))
}

func sum(i ...int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}
