package main

import "fmt"

type Person struct {
	first string
}

func changeName(p Person, name string) Person {
	p.first = name
	return p
}

func changeName2(p *Person, name string) {
	p.first = name
}

func main() {
	p1 := Person{
		first: "Aswin",
	}
	fmt.Println(p1.first)
	p1 = changeName(p1, "kumar")
	fmt.Println(p1.first)
	changeName2(&p1, "Aswin Kumar")
	fmt.Println(p1.first)

}
