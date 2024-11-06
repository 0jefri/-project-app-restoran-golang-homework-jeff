package utils

import (
	"fmt"
)

func DisplayError(message string, err error) {
	fmt.Printf("\033[31m%s %v\033[0m\n", message, err)
}

func DisplaySuccess(message string) {
	fmt.Printf("\033[32m%s\033[0m\n", message)
}
