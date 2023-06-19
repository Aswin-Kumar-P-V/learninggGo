package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	sum := sum(numbers)
	fmt.Println(sum)

}

func sum(numbers []int) (total int) {
	for _, val := range numbers {
		total += val
	}
	return
}
