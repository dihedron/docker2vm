package instruction

import (
	"regexp"

	log "github.com/dihedron/go-log"
)

// FLAVOUR is the pattern for the FLAVOUR instruction.
var FLAVOUR = regexp.MustCompile(`^\s*(?i)(flavour)(?-i)\s+([-a-zA-Z0-9\.]+)\s*$`)

// Flavour represents the FLAVOUR instruction in Packerfile format; this instruction
// does not exist in the Dockerfile schema.
type Flavour struct {
	// Token is the bare of the instruction, as returned by the Lexer.
	Token string
	// FlavourName is the name of the OpenStack Compute flavour to use for
	// launching the VM; even though the OpenStack Compute API has flavour IDs,
	// these have no defined format (they're just plain strings, and by default
	// are UUIDs), so they're indistinguishable from the name.
	FlavourName string
}

// newFlavour creates a new Flavour instruction and initilises it using
// information in the token via the associated regular expression.
func newFlavour(token string) ([]Instruction, error) {
	instruction := &Flavour{
		Token: token,
	}
	matches := FLAVOUR.FindStringSubmatch(instruction.Token)
	instruction.FlavourName = matches[2]
	log.Infof("FLAVOUR: using flavour: %q", instruction.FlavourName)
	return []Instruction{instruction}, nil
}
