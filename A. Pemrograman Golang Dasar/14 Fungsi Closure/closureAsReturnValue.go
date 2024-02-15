package main

import "fmt"

func main() {
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	var findMax = func(number []int, max int) (int, func() []int) {
		var result []int

		for _, e := range number {
			if e <= max {
				result = append(result, e)
			}
		}
		return len(result), func() []int {
			return result
		}
	}

	var max, find = findMax(numbers, 7)
	var getFind = find()
	fmt.Printf("Nilai Max : %v\nNilai Find : %v\n", max, getFind)
}
