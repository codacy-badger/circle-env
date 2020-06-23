package utils

import (
	"fmt"
	"os"
)

func Fatal(err error) {
	fmt.Println("Error:", err)
	os.Exit(1)
}
