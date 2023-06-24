package main

import "fmt"

type person struct {
	first string
}

func (p *person) speak() {
	fmt.Println("Hello I am ", p.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}
func main() {
	p := person{
		first: "Aswin",
	}
	saySomething(&p)
	p.speak()
}
