package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6}
	total := sum(xi...)
	fmt.Println("sum is:", total)

}
func sum(i ...int) int {
	fmt.Printf("i value: %v type : %v \n", i, i)
	var sum int
	for _, j := range i {
		sum += j
	}
	return sum
}
