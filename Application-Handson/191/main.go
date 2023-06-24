package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	p3 := person{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	people := []person{p1, p2, p3}

	fmt.Println(people)
	f, err1 := os.Create("./storeJson.txt")

	if err1 != nil {
		fmt.Println(err1)
	}

	err2 := json.NewEncoder(os.Stdout).Encode(people)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	err3 := json.NewEncoder(f).Encode(people)
	if err3 != nil {
		fmt.Println(err3)
		return
	}

}
