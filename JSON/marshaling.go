package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := Person{
		First: "Aswin",
		Last:  "Kumar",
		Age:   20,
	}
	p2 := Person{
		First: "Amal",
		Last:  "Krishna",
		Age:   22,
	}
	//This also works
	// persons := make([]Person, 0)
	// persons = append(persons, p1, p2)
	// fmt.Println(persons, len(persons))

	people := []Person{
		p1,
		p2,
	}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
