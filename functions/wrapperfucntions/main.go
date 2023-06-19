package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type content []byte

func (c content) String() string {
	return fmt.Sprint("Content of file:", string(c))
}

func main() {
	x := "test.txt"
	var err error
	var c content
	c, err = readFile(x)
	if err != nil {
		log.Fatalf("error encountered: %s", err)
	}
	fmt.Println(c)

}
func readFile(fileName string) ([]byte, error) {
	fp, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("file no opened due to: %s", err)
	}
	return fp, nil
}
