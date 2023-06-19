package main

import (
	"fmt"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("title of book:", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("No of books :", strconv.Itoa(int(c)))
}

func main() {
	b := book{
		title: "james bond",
	}

	var n count = 10

	fmt.Println(b)
	fmt.Println(n)
}
