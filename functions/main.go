package main

import "fmt"

func main() {
	foo()
	bar("Aswin")
	s := alias("Kumar")
	fmt.Println(s)
	name, age := nameAge("Aswin", 21)
	fmt.Printf("name : %v, age: %v \n", name, age)
}

// no params
func foo() {
	fmt.Println("iam foo")
}

func bar(s string) {
	fmt.Print("My name is:", s)
}
func alias(s string) string {
	return fmt.Sprint("Also known as :", s)
}
func nameAge(name string, age int) (string, int) {
	return fmt.Sprint("my name is:", name), age
}
