package main

import (
	"fmt"
	"strings"
)

func main() {
	var names [4]string

	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])

	var fruits [4]string

	fruits = [4]string{"apple", "grape", "banana", "melon"}
	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("isi semua element \t", fruits)

	fmt.Println(strings.Repeat("#", 10))
	fruits = [4]string{
		"apple",
		"grape",
		"banana",
		"watermelon",
	}
	fmt.Println("jumlah element \t\t=>", len(fruits))
	fmt.Println("isi semua element \t=>", fruits)

	fmt.Println(strings.Repeat("=>", 10))
	var numbers = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println("jumlah element numbers \t =>", len(numbers))
	fmt.Println("isi semua element \t =>", numbers)
}
