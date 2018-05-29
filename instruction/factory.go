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

func New(index int, token string) ([]Instruction, error) {
	if FROM.MatchString(token) {
		log.Debugf("found FROM instruction")
		return newFrom(token)
	} else if MAINTAINER.MatchString(token) {
		log.Debugf("found MAINTAINER instruction")
		return newMaintainer(token)
	} else if LABEL.MatchString(token) {
		log.Debugf("found LABEL instruction")
		return newLabel(token)
	} else if RUN.MatchString(token) {
		log.Debugf("found RUN instruction")
		return newRun(token)
	} else if ADD.MatchString(token) {
		log.Debugf("found ADD instruction")
		return newAdd(token)
	}
	return nil, fmt.Errorf("invalid instruction at line %d: %q", index, token)
}

var ()
