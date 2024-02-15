package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("angka", i)
	}

	// penggunaan for dengan argument hanya kondisi
	fmt.Println(strings.Repeat("#", 5))
	i := 0
	for i < 5 {
		fmt.Println("angka", i)
		i++
	}

	// for tanpa argument
	fmt.Println(strings.Repeat("#", 5))
	var index int = 0
	for {
		fmt.Println("angka", index)
		index++
		if index == 5 {
			break
		}
	}
}
