package main

import "fmt"

func fav(s string) string {
	return fmt.Sprint("My fav place is:", s)
}

func main() {
	fmt.Println(fav("Kozhikode"))
}
