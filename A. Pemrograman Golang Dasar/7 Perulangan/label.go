package main

import "fmt"

func main() {
outerloop:
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			if i == 3 {
				break outerloop
			}
			fmt.Printf("Matriks [%d] [%d]\n", i, j)
		}
	}
}
