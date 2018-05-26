package transpiler

import (
	"github.com/dihedron/docker2vm/instruction"
	log "github.com/dihedron/go-log"
)

type Parser struct {
	Instructions []instruction.Instruction
}

func (p *Parser) Parse(tokens []string) (string, error) {

	for index, token := range tokens {
		log.Debugf("parsing %q\n", token)
		instruction, err := instruction.New(index, token)
		if err != nil {
			log.Errorf("error: %v", err)
			continue
		}
		p.Instructions = append(p.Instructions, instruction)
	}
	return "", nil
}
