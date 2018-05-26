package main

import (
	"os"

	"github.com/dihedron/docker2vm/instruction"
	"github.com/dihedron/docker2vm/transpiler"
	"github.com/dihedron/go-log"
)

//
// docker2vm [Dockerfile]
//
// cat Dockerfile | docker2vm
//
func main() {

	log.SetLevel(log.DBG)
	log.SetStream(os.Stdout, true)
	log.SetTimeFormat("15:04:05.000")
	log.SetPrintCallerInfo(true)
	log.SetPrintSourceInfo(log.SourceInfoShort)

	var stream *os.File
	if len(os.Args) > 1 {
		var err error
		stream, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatalln(err)
		}
		defer stream.Close()
	} else {
		stream = os.Stdin
	}

	lexer := &transpiler.Lexer{}
	instructions, _ := lexer.Scan(stream)
	parser := &transpiler.Parser{
		Instructions: []instruction.Instruction{},
	}
	parser.Parse(instructions)
	// for i, instruction := range instructions {
	// 	fmt.Printf("[%3d] %s\n", i, instruction)
	// }
}
