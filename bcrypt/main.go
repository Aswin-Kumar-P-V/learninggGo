package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := `@sWin239`
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hashPass)
	loginPass := `@sWin239`
	err1 := bcrypt.CompareHashAndPassword(hashPass, []byte(loginPass))
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("Password matches")
	}
}
