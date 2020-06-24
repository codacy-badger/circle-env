package utils

import (
	"bufio"
	"os"
)

type IScanner interface {
	Scan() (string, error)
}

type scanner struct{}

var Scanner = new(scanner)

func (s *scanner) Scan() (string, error) {
	scn := bufio.NewScanner(os.Stdin)
	if !scn.Scan() {
		return "", scn.Err()
	}

	return scn.Text(), nil
}
