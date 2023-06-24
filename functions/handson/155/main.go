package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Println("hello iam ", p.first, "\n and iam of age", p.age)
}

func main() {

	p := person{
		first: "Aswin",
		age:   10,
	}
	p.speak()

}
