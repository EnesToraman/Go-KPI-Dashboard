package utils

import (
	"fmt"
	"os"
)

func CheckError(err error) {
	fmt.Println("Error: %w", err)
	os.Exit(1)
}
