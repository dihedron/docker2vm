package instruction

import (
	"regexp"

	log "github.com/dihedron/go-log"
)

// ADD is the pattern for the ADD instruction.
var ADD = regexp.MustCompile(`^\s*(?i)(?:add)(?-i)\s+(.*)\s+(.*)\s*$`)

// Add represents the ADD instruction, which is translated into the copy of a
// local file or diretory into the remove VM; source and destintaion must be
// specified.
type Add struct {
	// Token is the bare of the instruction, as returned by the Lexer.
	Token string
	// Source is the file or directory to copy.
	Source string
	// Destination is the destination path.
	Destination string
}

// newAdd creates a new Add instruction and initiliases it using information
// extracted from the token via the associated regular expression.
func newAdd(token string) ([]Instruction, error) {
	instruction := &Add{
		Token: token,
	}
	matches := ADD.FindStringSubmatch(instruction.Token)
	instruction.Source = matches[1]
	instruction.Destination = matches[2]
	log.Infof("ADD: adding add: %q => %q", instruction.Source, instruction.Destination)
	return []Instruction{instruction}, nil
}
