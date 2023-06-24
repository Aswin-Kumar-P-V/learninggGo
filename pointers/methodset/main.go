package main

import "fmt"

type Dog struct {
	first string
}

func (d Dog) Barks() {
	fmt.Println("My name is:", d.first, "And i love to bark")
}
func (d *Dog) Walks() {
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
	d1.Barks()
	d1.Walks()
	d2 := &Dog{
		first: "potter",
	}
	//x := Young(d1)
	//Throws an error

	// d2.Barks()
	// d2.Walks()
	x := Young(d2)
	x.Barks()
	x.Walks()
}
