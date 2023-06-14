package main

import "fmt"

func main() {
	// a := [2]int{0, 1}
	// fmt.Printf("Array a is %v \n", a)
	// b := [...]string{"Hello", "world"}
	// fmt.Printf("Array b is %v \n", b)

	// var c [2]int
	// c[0] = 1
	// c[1] = 2
	// fmt.Printf("Array c is %v\n", c)

	// a := []int{1, 2, 3, 4}
	// fmt.Printf("a : %v is %T and length is %v", a, a, len(a))
	// for i, v := range a {
	// 	fmt.Printf("index : %v val : %v", i, v)
	// }

	// a := []int{1, 2, 3, 4}
	// fmt.Printf("a : %v is %T and length is %v \n", a, a, len(a))
	// for i, _ := range a {
	// 	fmt.Printf("index : %v val : %v \n", i, a[i])
	// }

	// a := []int{1, 2, 3, 4}
	// fmt.Println(a)

	// a = append(a, 5, 6, 7)
	// fmt.Println(a)

	// b := []int{8, 9, 10}

	// a = append(a, b...)

	// fmt.Println(a[:])

	// a = append(a[:3], a[4:]...)

	// fmt.Println(a[:])

	// a := make([]int, 0, 5)
	// fmt.Printf("a has len:%v and capacity:%v ", len(a), cap(a))
	// a = append(a, 1, 2, 3, 4, 5, 6)
	// fmt.Printf("a has len:%v and capacity:%v ", len(a), cap(a))

	// person1 := []string{"aswin", "kumar", "21"}
	// person2 := []string{"Dakshina", "D", "20"}
	// persons := [][]string{person1, person2}
	// fmt.Println(persons)
	// fmt.Println(persons[0][0])

	// a := []int{1, 2, 3, 4}
	// b := a

	// fmt.Println(a, b)

	// a[0] = 5

	// fmt.Println(a, b)

	c := []int{5, 6, 7, 9}
	d := make([]int, 5)
	copy(d, c)
	fmt.Println(c, d)
	c[0] = 1
	fmt.Println(c, d)

}
