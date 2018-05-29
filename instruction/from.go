package instruction

import (
	"regexp"
	"strings"

	log "github.com/dihedron/go-log"
)

// FROM is the pattern for the FROM instruction; it may optionally include the
// VM flavour in the AS clause.
//([-a-zA-Z0-9\.]+)
var FROM = regexp.MustCompile(`^\s*(?i)(?:from)(?-i)\s+([-a-zA-Z0-9]+)(?:\s+(?:(?i)(?:as)(?-i)\s+([-a-zA-Z0-9\.]+))){0,1}\s*$`)

// From represents the FROM instruction in Packerfile/Dockerfile format.
type From struct {
	// Token is the bare of the instruction, as returned by the Lexer.
	Token string
	// ImageName is the name of the OpenStack image to use for creating the VM.
	ImageName *string
	// ImageID is the ID of the OpenStack image to use for creating the VM.
	ImageID *string
	// FlavourName is the name of the OpenStack Compute flavour to use for
	// launching the VM; even though the OpenStack Compute API has flavour IDs,
	// these have no defined format (they're just plain strings, and by default
	// are UUIDs), so they're indistinguishable from the name.
	FlavourName string
}

// newFrom creates a new From instruction and initilises it using information in
// the token via the associated regular expression.
func newFrom(token string) ([]Instruction, error) {
	instruction := &From{
		Token: token,
	}
	matches := FROM.FindStringSubmatch(instruction.Token)
	if len(matches) > 2 && len(strings.TrimSpace(matches[2])) > 0 {
		instruction.FlavourName = matches[2]
	}
	re := regexp.MustCompile(`^\s*(?i)([0-9a-fA-f]{8}-[0-9a-fA-f]{4}-[0-9a-fA-f]{4}-[0-9a-fA-f]{4}-[0-9a-fA-f]{12})(?-i)\s*$`)
	if re.MatchString(matches[1]) {
		instruction.ImageID = &matches[1]
		log.Infof("FROM: using image id: %q and VM flavour: %q", *instruction.ImageID, instruction.FlavourName)
	} else {
		instruction.ImageName = &matches[1]
		log.Infof("FROM: using image name: %q and VM flavour: %q", *instruction.ImageName, instruction.FlavourName)
	}
	return []Instruction{instruction}, nil
}
