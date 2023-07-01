package main

import (
	"fmt"

	"github.com/Aswin-Kumar-P-V/learninggGo/Documentation/dog"
)

type doberman struct {
	name string
	age  int
}

func (d doberman) String() string {
	return fmt.Sprintf("Dog's name is %v \nand Age is %v \n", d.name, d.age)
}

func main() {
	roksy := doberman{
		name: "Roksy",
		age:  dog.Years(3),
	}
	fmt.Println(roksy)
}
