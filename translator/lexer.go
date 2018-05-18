package translator

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

// Lexer is the type representing a Dockerfile analyser capable of extracting
// tokens from the inpu file or stream; it skips comments and reconstructs
// instructions even when continued on multiple lines.
type Lexer struct {
}

// Tokenise scans a Dockerfile line by line, assembling continued lines as
// instructions.
func (l Lexer) Tokenise(r io.Reader) ([]string, error) {
	lines := []string{}
	scanner := bufio.NewScanner(r)
	no := 1
	var buffer bytes.Buffer
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "#") {
			// this is a comment, skip it
		} else if len(buffer.String()) > 0 {
			// we are in the middle of a continuation
			if strings.HasSuffix(scanner.Text(), "\\") {
				// this is a continuation, append to the buffer
				text := scanner.Text()[:len(scanner.Text())-len("\\")]
				buffer.WriteString(text)
			} else {
				// this ends a buffer
				if len(strings.TrimSpace(scanner.Text())) > 0 {
					// only append if it's not an empty line
					buffer.WriteString(scanner.Text())
				}
				lines = append(lines, buffer.String())
				buffer.Reset()
			}
		} else {
			// no continuation
			if strings.HasSuffix(scanner.Text(), "\\") {
				// this is the beginning of a continuation, append to the buffer
				text := scanner.Text()[:len(scanner.Text())-len("\\")]
				buffer.WriteString(text)
			} else {
				// this line is by itself
				if len(strings.TrimSpace(scanner.Text())) > 0 {
					// append it only if not empty
					lines = append(lines, scanner.Text())
				}
			}
		}
		no++
	}
	if len(buffer.String()) > 0 {
		// the last continuation was not closed: this is an error
		return nil, fmt.Errorf("error at end of file: unterminated instruction")
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
