package utils

import (
	"bufio"
	"os"
)

type IScanner interface {
	Scan() (string, error)
}

type Scanner struct{}

func NewScanner() *Scanner {
	return new(Scanner)
}

func (s *Scanner) Scan() (string, error) {
	scn := bufio.NewScanner(os.Stdin)
	if !scn.Scan() {
		panic(scn.Err())
	}

	return scn.Text(), nil
}
