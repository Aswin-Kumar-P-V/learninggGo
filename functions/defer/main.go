package main

import "fmt"

func main() {
	defer one()
	two()

}
func one() {
	fmt.Println("one")
}
func two() {
	fmt.Println("two")
}
