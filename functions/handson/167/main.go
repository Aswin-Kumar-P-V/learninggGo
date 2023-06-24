package main

import (
	"fmt"
	"time"
)

func main() {
	timeFunc(doWork)

}

func doWork() {
	for i := 0; i < 2_000; i++ {
		fmt.Println("HELLO")
	}
}

func timeFunc(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
