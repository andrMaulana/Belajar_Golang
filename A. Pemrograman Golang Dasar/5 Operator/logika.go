package main

import "fmt"

func main() {
	var right bool = true
	var left bool = false

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t)\n", leftAndRight)
	var leftOrRight = left || right
	fmt.Printf("left or right \t(%t)\n", leftOrRight)
	var leftReverse = !left
	fmt.Printf("not left \t(%t)\n", leftReverse)
}
