package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"Aswin","Last":"Kumar","Age":20},{"First":"Amal","Last":"Krishna","Age":22}]`
	bs := []byte(s)
	fmt.Printf("%T:\n", s)
	fmt.Printf("%T:\n", bs)
	//This also works
	// people := []person{}
	// fmt.Printf("%T:\n", people)

	var people []person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)

}
