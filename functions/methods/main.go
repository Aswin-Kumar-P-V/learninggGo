package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("my name is :", p.first)
}

func main() {
	p1 := person{
		first: "aswin",
	}
	p2 := person{
		first: "daksina",
	}
	p1.speak()
	p2.speak()

}
