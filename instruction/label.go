package instruction

import (
	"regexp"

	log "github.com/dihedron/go-log"
)

// MAINTAINER is the pattern for the MAINTAINER instruction, which is regarded as
// a special (deprecated) case of a Label, with a fixed key "maintainer".
//var MAINTAINER = regexp.MustCompile(`^\s*(?i)(maintainer)(?-i)\s+([a-zA-Z0-9\._@<>\(\)]+)\s*$`)
var MAINTAINER = regexp.MustCompile(`^\s*(?i)(maintainer)(?-i)\s+(.*)\s*$`)

// LABEL is the pattern for the LABEL instruction, which is a flexible way to
// store tags in a VM image.
//var LABEL = regexp.MustCompile(`^\s*(?i)(label)(?-i)\s+(.*)\s*$`)
var LABEL = regexp.MustCompile(`^\s*(?i)(?:label)(?-i)\s+(?:(?:"{0,1})([^"]*)(?:"{0,1})="([^"]*)")+\s*$`)

//var LABELTAGS = regexp.MustCompile(`(.*)="(.*)"`)

// Label represents the LABEL and MAINTAINER instructions in Packerfile format;
// the MAINTAINER instruction is still quite commonly used in Dockerfiles but
// it has been deprecated for some time and it should be replaced with the
// more general LABEL instruction. The Parser will issue a warning and treat it
// as a Label instruction with a fixed "maintainer" key.
type Label struct {
	// Token is the bare of the instruction, as returned by the Lexer.
	Token string
	// Value is a the value associated with the Label.
	Key string
	// Value is a the value associated with the Label.
	Value string
}

// newMaintainer creates a new Label instruction and initilises it using
// information extracted from the token via the associated regular expression.
func newMaintainer(token string) ([]Instruction, error) {
	instruction := &Label{
		Token: token,
		Key:   "maintainer",
	}
	matches := MAINTAINER.FindStringSubmatch(instruction.Token)
	instruction.Value = matches[2]
	log.Warnf("MAINTAINER: using deprecated maintainer: %q => %q", instruction.Key, instruction.Value)
	return []Instruction{instruction}, nil
}

// newLabel creates one or more new Label instructions and initilises them using
// information extracteed from the token via the associated regular expression.
func newLabel(token string) ([]Instruction, error) {
	instructions := []Instruction{}
	matches := LABEL.FindStringSubmatch(token)
	for i, match := range matches {
		log.Infof("match[%d] => %q", i, match)
	}
	/*
		log.Infof("match: %q", matches[2])
		//pairs := LABELTAGS.FindAllStringSubmatch(matches[2], -1)
		for _, pair := range pairs {
			instruction := &Label{
				Token: token,
			}
			log.Infof(" - submatch: %q => %q", pair[0], pair[1])
			instructions = append(instructions, instruction)
		}
		//i.Value = matches[2]
		//log.Infof("LABEL: using deprecated maintainer: %q => %q", i.Key, i.Value)
	*/
	return instructions, nil
}
