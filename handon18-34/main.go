package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where initialization starts")
}

func main() {
	// x := rand.Intn(250)
	// fmt.Printf("value of x: %v \n", x)
	//using if

	// if x < 100 {
	// 	fmt.Println("x is less than 100")
	// } else if x < 200 {
	// 	fmt.Println("x is btn 100 and 200")
	// } else {
	// 	fmt.Println("x is greater than 200")
	// }

	//using switch
	/*
		switch {
		case x < 100:
			fmt.Println("x is less than 100")
		case x < 200:
			fmt.Println("x is btn 100 and 200")
		default:
			fmt.Println("x is greater than 200")
		}
	*/

	// for i := 0; i < 100; i++ {
	// 	fmt.Printf("i : %v \n", i)
	// }

	// for i := 0; i < 43; i++ {
	// 	a := rand.Intn(6)
	// 	switch a {
	// 	case 1:
	// 		fmt.Printf("a is %v \n", 1)
	// 	case 2:
	// 		fmt.Printf("a is %v \n", 2)
	// 	case 3:
	// 		fmt.Printf("a is %v \n", 3)
	// 	case 4:
	// 		fmt.Printf("a is %v \n", 4)
	// 	default:
	// 		fmt.Printf("a is %v \n", 5)
	// 	}
	// }

	// for i := 0; i < 5; i++ {
	// 	for j := 0; j < 5; j++ {
	// 		fmt.Println("j \t")
	// 	}
	// 	fmt.Println("\n")
	// }

	// xi := []int{1, 2, 3, 4, 5}

	// for i, val := range xi {
	// 	fmt.Printf("index : %v, value : %v\n", i, val)
	// }

	// m := map[string]int{
	// 	"james": 42,
	// 	"aswin": 12,
	// }

	// for key, val := range m {
	// 	fmt.Printf("key : %v, value : %v \n", key, val)
	// }
	// if q, ok := m["fafa"]; !ok {
	// 	fmt.Printf("no value only %v \n", q)
	// }

	// if age, ok := m["james"]; ok {
	// 	fmt.Printf("yes , age of james is %v", age)
	// }
	for i := 0; i < 20; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("value pf x is 3")
		}
	}
}
