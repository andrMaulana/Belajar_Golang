package main

import (
	"fmt"
	"math/rand"
	"time"
)

var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

const overtime = 40

func main() {
	var randomValue int

	randomValue = randomWithRange(3, 10)
	fmt.Println(randomValue)
	randomValue = randomWithRange(4, 6)
	fmt.Println(randomValue)

	divideNumber(10, 14)
	divideNumber(12, 3)
	fmt.Println("number of layer =>", PreparationTime(1, 30))
}
func randomWithRange(min, max int) int {
	var value = randomizer.Int()%(min-max+1) + min
	return value
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Println("error!. cannot divide 0")
		return
	}
	divide := m / n
	fmt.Println(divide)
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers, actualMinutesInOven int) int {
	return numberOfLayers*2 + actualMinutesInOven
}
