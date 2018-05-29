package instruction

import (
	"regexp"

	log "github.com/dihedron/go-log"
)

// MAINTAINER is the pattern for the MAINTAINER instruction, which is regarded as
// a special (deprecated) case of a Label, with a fixed key "maintainer".
var MAINTAINER = regexp.MustCompile(`^\s*(?i)(?:maintainer)(?-i)\s+(.*)\s*$`)

// LABEL is the pattern for the LABEL instruction, which is a flexible way to
// store tags in a VM image.
var LABEL = regexp.MustCompile(`^\s*(?i)(?:label)(?-i)\s+(.*)\s*$`)

// ITEM represents a single key/value pair (a label) in a LABEL instruction.
var ITEM = regexp.MustCompile(`\s*(?:(?:"{0,1})([^"]*)(?:"{0,1})="([^"]*)")`)

// Label represents the LABEL and MAINTAINER instructions in Packerfile format;
// the MAINTAINER instruction is still quite commonly used in Dockerfiles but
// it has been deprecated for some time and it should be replaced with the
// more general LABEL instruction. The Parser will issue a warning and treat it
// as a Label instruction with a fixed "maintainer" key.
type Label struct {
	// Token is the bare of the instruction, as returned by the Lexer.
	Token string
	// Key is the name of the Label.
	Key string
	// Value is the value associated with the Label.
	Value string
}

// newMaintainer creates a new Label instruction and initialises it using
// information extracted from the token via the associated regular expression.
func newMaintainer(token string) ([]Instruction, error) {
	instruction := &Label{
		Token: token,
		Key:   "maintainer",
	}
	matches := MAINTAINER.FindStringSubmatch(instruction.Token)
	instruction.Value = matches[1]
	log.Warnf("MAINTAINER: using deprecated maintainer: %q => %q", instruction.Key, instruction.Value)
	return []Instruction{instruction}, nil
}

// newLabel creates one or more new Label instructions and initilises them using
// information extracteed from the token via the associated regular expression.
func newLabel(token string) ([]Instruction, error) {
	instructions := []Instruction{}
	matches := LABEL.FindStringSubmatch(token)
	labels := ITEM.FindAllStringSubmatch(matches[1], -1)
	for _, label := range labels {
		instruction := &Label{
			Token: token,
			Key:   label[1],
			Value: label[2],
		}
		instructions = append(instructions, instruction)
		log.Infof("LABEL: adding label: %q => %q", instruction.Key, instruction.Value)
	}
	return instructions, nil
}
