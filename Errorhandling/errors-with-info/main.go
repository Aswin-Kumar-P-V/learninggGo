package main

import (
	"errors"
	"fmt"
	"math"
)

var negSquareErr = errors.New("Square root of a negative number not available")

func main() {
	ans, err := sqrt(25)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ans)

	ans, err = sqrt(-10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ans)

}

func sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, negSquareErr
	} else {
		return math.Pow(num, 0.5), nil
	}
}
