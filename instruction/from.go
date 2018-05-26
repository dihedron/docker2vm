package instruction

import (
	"regexp"

	log "github.com/dihedron/go-log"
)

// FROM is the pattern for the FROM instruction.
var FROM = regexp.MustCompile(`^\s*(?i)(from)(?-i)\s+([-a-zA-Z0-9]+)\s*$`)

// From represents the FROM instruction in Packerfile/Dockerfile format.
type From struct {
	// Token is the bare of the instruction, as returned by the Lexer.
	Token string
	// ImageName is the name of the OpenStack image to use for creating the VM.
	ImageName *string
	// ImageID is the ID of the OpenStack image to use for creating the VM.
	ImageID *string
}

func (f *From) Init(instruction string) (Instruction, error) {
	f.Token = instruction
	matches := FROM.FindStringSubmatch(f.Token)
	re := regexp.MustCompile(`^\s*(?i)([0-9a-fA-f]{8}-[0-9a-fA-f]{4}-[0-9a-fA-f]{4}-[0-9a-fA-f]{4}-[0-9a-fA-f]{12})(?-i)\s*$`)
	if re.MatchString(matches[2]) {
		f.ImageID = &matches[2]
		log.Infof("FROM: using image id: %q", *f.ImageID)
	} else {
		f.ImageName = &matches[2]
		log.Infof("FROM: using image name: %q", *f.ImageName)
	}
	return f, nil
}
