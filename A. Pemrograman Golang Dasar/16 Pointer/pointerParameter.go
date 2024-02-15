package main

import "fmt"

func main() {
	var original = 4

	changeNUmber(&original, 10)
	fmt.Println(original)
}

func changeNUmber(number *int, value int)  {
	*number = value
}
