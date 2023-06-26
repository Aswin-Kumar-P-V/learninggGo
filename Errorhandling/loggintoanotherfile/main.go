package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("Error.txt")
	if err != nil {
		fmt.Println("Error happened while creating Error.txt file", err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err1 := os.Open("New.txt")
	if err1 != nil {
		log.Println("Error happended while opening New.txt file", err1)
	}
	defer f2.Close()
	fmt.Println("exiting")
}
