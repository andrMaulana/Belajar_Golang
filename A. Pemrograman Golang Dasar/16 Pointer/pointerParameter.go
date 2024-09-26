package main

import "fmt"

func main() {
	var number int = 4

	fmt.Println("Number before (value) :", number)
	fmt.Println("Number before (Address) :", &number)

	Change(&number, 10)
	fmt.Println("Number after (value) :", number)
	fmt.Println("Number after (Address) :", &number)
}

func Change(original *int, number int) {
	*original = number
}
