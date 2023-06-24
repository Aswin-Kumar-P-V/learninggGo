package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Architecture : ", runtime.GOARCH)
	fmt.Println("OS : ", runtime.GOOS)
}
