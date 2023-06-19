package main

import (
	"fmt"
	"log"
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
func loginfo(s fmt.Stringer) {
	log.Println("log from logger", s.String())
}
func main() {
	b := book{
		title: "james bond",
	}

	var n count = 10

	loginfo(b)
	loginfo(n)
}
