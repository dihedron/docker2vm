package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/dihedron/docker2vm/parser"
)

//
// docker2vm [file]
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

	parser := &parser.Parser{}
	instructions, _ := parser.Parse(reader)
	for i, instruction := range instructions {
		fmt.Printf("[%3d] %s\n", i, instruction)
	}

}
