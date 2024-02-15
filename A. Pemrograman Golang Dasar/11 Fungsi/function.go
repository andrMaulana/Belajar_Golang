package main

import (
	"fmt"
	"strings"
)

func main() {
	var names = []string{"john", "wick"}

	printMessages("hallo", names)
}

func printMessages(message string, arr []string) {
	joinString := strings.Join(arr, " ")
	fmt.Println(message, joinString)
}
