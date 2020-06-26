package utils

import (
	"fmt"
)

// Confirm ...
func Confirm(message string) (bool, error) {
	fmt.Printf("%s", message)
	ipt, err := Scanner.Scan()
	if err != nil {
		return false, err
	}

	return ipt == "yes", nil
}
