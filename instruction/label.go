package instruction

import (
	"regexp"

	log "github.com/dihedron/go-log"
)

// MAINTAINER is the pattern for the MAINTAINER instruction, which is regarded as
// a special (deprecated) case of a Label, with a fixed key "maintainer".
//var MAINTAINER = regexp.MustCompile(`^\s*(?i)(maintainer)(?-i)\s+([a-zA-Z0-9\._@<>\(\)]+)\s*$`)
var MAINTAINER = regexp.MustCompile(`^\s*(?i)(maintainer)(?-i)\s+(.*)\s*$`)

// LABEL is the pattern for the LABEL instrcution, which is a flexible way to
// store tags in a VM image.
var LABEL = regexp.MustCompile(`^\s*(?i)(label)(?-i)\s+(.*)\s*$`)

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
// information in the token via the associated regular expression.
func newMaintainer(instruction string) (Instruction, error) {
	i := &Label{
		Token: instruction,
		Key:   "maintainer",
	}
	matches := MAINTAINER.FindStringSubmatch(i.Token)
	i.Value = matches[2]
	log.Warnf("MAINTAINER: using deprecated maintainer: %q => %q", i.Key, i.Value)
	return i, nil
}
