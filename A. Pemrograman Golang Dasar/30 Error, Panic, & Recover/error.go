package main

import (
	"errors"
	"fmt"
	"strings"
)

func ValidateError(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("not empty")
	}
	return true, nil
}
func main() {
	var input string
	fmt.Scanln(&input)

	if valid, err := ValidateError(input); valid {
		fmt.Println(input)
	} else {
		fmt.Println(err.Error())
	}
}
