package main

import "fmt"

func main() {

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var newNumber = func(min int) []int {
		var result []int

		for _, e := range numbers {
			if e < min {
				continue
			}
			result = append(result, e)
		}

		return result
	}(4)

	var min, max = func(n []int) (int, int) {
		var min, max int

		for i, e := range n {
			switch {
			case i == 0:
				min, max = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}(numbers)
	fmt.Println(newNumber)
	fmt.Printf("Nilai Numbers : %v\nNilai Min : %v\nNilai Max : %v\n", numbers, min, max)
}
