package main

import (
	"fmt"
	"strings"
)

type FuncCallback func(string) bool

func filter(data []string, callback FuncCallback) []string {
	var result []string

	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}

	return result
}

func main() {
	var word = []string{"done", "win", "tools", "gip", "jab"}
	var stringO = filter(word, func(s string) bool {
		return strings.Contains(s, "o")
	})

	fmt.Println(stringO) // sould must be : ["done", "tools"]
}
