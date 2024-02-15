package main

import "fmt"

func main() {
	// gaya penulisan secara horizontal
	var chicken = map[string]int{"januari": 50, "februari": 40}

	// penulisan secara vertical
	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}
	fmt.Println(chicken["januari"])
	fmt.Println(chicken["februari"])

	fmt.Println(chicken2["januari"])
	fmt.Println(chicken2["februari"])

}
