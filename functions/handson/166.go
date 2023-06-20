package main

import (
	"fmt"
	"math"
)

func pow(a float64) func() float64 {
	var c float64
	return func() float64 {
		c++
		return math.Pow(a, c)
	}
}

func main() {
	x := pow(10)
	y := x()
	fmt.Println(y)
	y = x()
	fmt.Println(y)
	y = x()
	fmt.Println(y)
}
