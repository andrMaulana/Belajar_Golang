package main

import (
	"errors"
	"fmt"
	"strings"
)

func ValidateError(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Not empty")
	}

	return true, nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Terjadi Error :", r)
	} else {
		fmt.Println("Aplikasi Berjalan Sempurna")
	}
}
func main() {
	defer catch()
	var name string
	fmt.Print("Tuliskan Teks disini ")
	fmt.Scanln(&name)

	if valid, err := ValidateError(name); valid {
		fmt.Println("Hallo", name)

	} else {
		panic(err.Error())
	}
}
