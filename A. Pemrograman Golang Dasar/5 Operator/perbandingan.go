package main

import "fmt"

func main() {
	var aritmatika = (((2+6)%3)*4 - 2) / 3
	var isEquals = (aritmatika == 2)
	fmt.Printf("nilai %d (%t) \n", aritmatika, isEquals)
}
