package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}
func main() {
	runtime.GOMAXPROCS(2)

	print(5, "apa kabar?")
	go print(5, "hallo")

	var input string
	fmt.Scanln(&input)
}
