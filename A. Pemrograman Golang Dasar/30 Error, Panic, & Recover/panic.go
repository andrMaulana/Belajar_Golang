package main

import (
	"errors"
	"fmt"
	"strings"
)

func ValidateError(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Not Empty")

	}
	return true, nil
}
func main() {
	var name string
	fmt.Print("Type Your Name ..")
	fmt.Scanln(&name)

	if valid, err := ValidateError(name); valid {
		fmt.Println("Hallo", name)

	} else {
		panic(err.Error())
		fmt.Println("end")
	}
}
