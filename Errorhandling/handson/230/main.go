package main

import "fmt"

type customErr struct {
	info string
}

func (cE customErr) Error() string {
	return fmt.Sprintf("The error is : %v", cE.info)
}

func main() {
	c1 := customErr{
		info: "Consistency is missing, need to be more discplined",
	}
	fmt.Println(c1)
	foo(c1)

}

func foo(e error) {
	fmt.Println("hello from foo")
	fmt.Println("And the error is:", e)
	//Here using e.info wont work instead we have to use e.(customErr).info this is called assertion (not conversion)
	fmt.Println("The error is :", e.(customErr).info)
}
