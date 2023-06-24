package main

import "fmt"

func sumVar(a ...int) int {
	total := 0
	for _, val := range a {
		total += val
	}
	return total
}

func main() {
	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(sumVar(xi...))
}
