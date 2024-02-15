package main

import "fmt"

func main() {
	var getMinMax = func(number ...int) (int, int) {
		var min, max int

		for i, e := range number {
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
	}
	var number = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(number...)
	fmt.Printf("Nilai numbers : %v\nNilai Min : %v\nNilai Max : %v\n", number, min, max)
}
