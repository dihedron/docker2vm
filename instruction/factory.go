package instruction

import (
	"fmt"

	log "github.com/dihedron/go-log"
)

// Instruction is the common interface to all valid Packerfile instructions.
type Instruction interface {
	// Init inisalises an instruction; it may issue queries such as retrieving
	// information by performing a GET request to the cloud provider.
	//Init(instruction string) (Instruction, error)
	// more methods here
}

func New(index int, instruction string) (Instruction, error) {
	if FROM.MatchString(instruction) {
		log.Debugf("found FROM instruction")
		return newFrom(instruction)
	} else if FLAVOUR.MatchString(instruction) {
		log.Debugf("found FLAVOUR instruction")
		return newFlavour(instruction)
	} else if MAINTAINER.MatchString(instruction) {
		log.Debugf("found MAINTAINER instruction")
		return newMaintainer(instruction)
	}
	return nil, fmt.Errorf("invalid instruction at line %d: %q", index, instruction)
}

var ()
