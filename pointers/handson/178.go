package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}
func main() {
	fmt.Printf("a Value: %v Type: %t \n", a, a)
	fmt.Printf("b Value: %v Type: %t \n", b, b)
	fmt.Printf("c Value: %v Type: %t \n", c, c)
	fmt.Printf("d Value: %v Type: %t \n", d, d)
	fmt.Println("-----CONTENT-------")
	fmt.Println(*a)
	fmt.Println(*b)
	fmt.Println(*c)
	fmt.Println(*d)

}
