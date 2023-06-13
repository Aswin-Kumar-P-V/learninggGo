package main

import (
	"fmt"
	"math/rand"
)

var b = 2

const c = 3

func main() {
	a := 42

	switch {
	case a < 42:
		fmt.Println("a is less")
	case a == 42:
		fmt.Println("a is equal")
	default:
		fmt.Println("hello")
	}

	options := 1
	switch options {
	case 1:
		fmt.Println("Option 1")
	case 2:
		fmt.Println("Option 2")
	}

	if z := 2 * rand.Intn(100); z < a {
		fmt.Printf("z is lesser %v \n", z)
	} else if z == a {
		fmt.Printf("z is equal %v \n", z)
	} else {
		fmt.Printf("Z is greater %v \n", z)
	}

	i := -1
	for i < 5 {
		fmt.Printf("i:%v\n", i)
		i++
	}
}
