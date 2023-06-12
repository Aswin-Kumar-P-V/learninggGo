package main

import (
	"fmt"

	"github.com/Aswin-Kumar-P-V/bird"
)

func main() {
	s1 := bird.Chirp()
	s2 := bird.Chirps()

	fmt.Println(s1, " ", s2)
	fmt.Printf("No of birds chirping is %d", bird.BirdsNo())

}
