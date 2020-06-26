package utils

import (
	"bufio"
	"os"
	"strings"
)

// IScanner ...
type IScanner interface {
	Scan() (string, error)
}

type scanner struct{}

// Scanner ...
var Scanner = new(scanner)

func (s *scanner) Scan() (string, error) {
	scn := bufio.NewScanner(os.Stdin)
	if !scn.Scan() {
		return "", scn.Err()
	}

	return strings.TrimSpace(scn.Text()), nil
}
