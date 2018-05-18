package main

import (
	"fmt"
	"io"
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

	var reader io.Reader
	if len(os.Args) > 1 {
		reader, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer reader.Close()
	} else {
		reader = os.Stdin
	}

	lexer := &translator.Lexer{}
	instructions, _ := lexer.Tokenise(reader)
	for i, instruction := range instructions {
		fmt.Printf("[%3d] %s\n", i, instruction)
	}
}
