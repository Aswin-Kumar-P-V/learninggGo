package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) writeInto(w io.Writer) error {
	_, err := w.Write([]byte(p.first))
	return err
}

func main() {
	p := person{
		first: "Aswin kumar",
	}
	f, err := os.Create("test.txt")
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	defer f.Close()
	// s := []byte("Hello world")
	// _, err1 := f.Write(s)
	// if err1 != nil {
	// 	log.Fatalf("ERROR : %s", err1)
	// }
	// b := bytes.NewBufferString("hello")
	// fmt.Println(b.String())
	// b.WriteString("World")
	// fmt.Println(b.String())
	// b.Reset()
	// b.Write([]byte("Hello gophers"))
	// fmt.Println(b.String())
	var b bytes.Buffer
	p.writeInto(f)
	p.writeInto(&b)
	fmt.Println(b.String())

}
