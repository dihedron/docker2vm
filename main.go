package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dihedron/docker2vm/translator"
)

//
// docker2vm [Dockerfile]
//
// cat Dockerfile | docker2vm
//
func main() {

	var stream *os.File
	if len(os.Args) > 1 {
		var err error
		stream, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer stream.Close()
	} else {
		stream = os.Stdin
	}

	reader := &translator.Reader{}
	instructions, _ := reader.Read(stream)
	for i, instruction := range instructions {
		fmt.Printf("[%3d] %s\n", i, instruction)
	}
}
