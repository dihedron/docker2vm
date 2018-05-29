package instruction

import (
	"regexp"

	log "github.com/dihedron/go-log"
)

// RUN is the pattern for the RUN instruction.
var RUN = regexp.MustCompile(`^\s*(?i)(?:run)(?-i)\s+(.*)\s*$`)

// Run represents the RUN instruction, which is translated into the execution of
// a specific instruction by a shell provisioner.
type Run struct {
	// Token is the bare of the instruction, as returned by the Lexer.
	Token string
	// Command is the command to run.
	Command string
}

// newRun creates a new Run instruction and initiliases it using
// information extracted from the token via the associated regular expression.
func newRun(token string) ([]Instruction, error) {
	instruction := &Run{
		Token: token,
	}
	matches := RUN.FindStringSubmatch(instruction.Token)
	instruction.Command = matches[1]
	log.Infof("RUN: adding run: %q", instruction.Command)
	return []Instruction{instruction}, nil
}
