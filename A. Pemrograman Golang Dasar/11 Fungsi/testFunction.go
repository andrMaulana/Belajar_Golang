package main

import "fmt"

func main() {
	knightIsAwake := true
	archerIsAwake := true
	prisonerIsAwake := true
	petDogIsPresent := true

	fmt.Println(CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent))
}

func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) (bool, string) {
	if knightIsAwake && archerIsAwake {
		return false, "kondisi satu"
	} else if prisonerIsAwake || petDogIsPresent {
		return false, "kondisi 2"
	} else {
		return true, "kondisi 4"
	}
}
