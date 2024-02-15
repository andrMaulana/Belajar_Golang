package main

import (
	"fmt"
	"strings"
)

func main() {
	var numbers = calculated(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var avg = fmt.Sprintf("nilai rata-rata = %.2f\n", numbers)
	fmt.Println(avg)

	fmt.Println(strings.Repeat("=", 10))

	var number2 = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	var avg1 = calculated(number2...)
	fmt.Printf("Nilai rata-rata : %.2f\n", avg1)
}

func calculated(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	var avg = float64(total) / float64(len(numbers))
	return avg
}
