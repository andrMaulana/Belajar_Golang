package main

import "fmt"

func main() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("Nilai Variable Number A = ", numberA)
	fmt.Println("Nilai Pointer Variable A = ", &numberA)

	fmt.Println("Nilai Variable Number B = ", *numberB)
	fmt.Println("Nilai Pointer Number B = ", numberB)

	numberA = 5
	fmt.Println("Nilai Number A", numberA)
	fmt.Println("Nilai Number B", *numberB)
}
