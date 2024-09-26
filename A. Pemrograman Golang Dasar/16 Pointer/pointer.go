package main

import "fmt"

func main() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("number A ( value ) :", numberA)
	fmt.Println("number A (Address) :", &numberA)

	fmt.Println("number B ( value ) :", *numberB)
	fmt.Println("number B ( Address ) :", numberB)

	// Efek perubahan nilai pointer
	fmt.Println("")

	numberA = 5

	fmt.Println("Number A ( Value ) :", numberA)
	fmt.Println("Number A ( Address ) :", &numberA)

	fmt.Println("Number B ( Value ) :", *numberB)
	fmt.Println("Number B ( Address ) :", numberB)
}
