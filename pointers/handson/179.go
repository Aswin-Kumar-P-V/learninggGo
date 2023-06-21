package main

import "fmt"

type Dog struct {
	first string
}

func (d Dog) Barks() {
	fmt.Println("My name is:", d.first, "And i love to bark")
}
func (d Dog) Walks() {
	d.first = "Barbara"
	fmt.Println("My name is:", d.first, "And i love to Walk")
}

type Young interface {
	Barks()
	Walks()
}

func main() {
	d1 := Dog{
		first: "Harry",
	}

	d2 := &Dog{
		first: "potter",
	}
	x := Young(d1)
	x.Barks()
	x.Walks()
	y := Young(d2)
	y.Barks()
	y.Walks()
}
