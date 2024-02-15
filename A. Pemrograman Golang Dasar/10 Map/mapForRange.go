package main

import (
	"fmt"
)

func main() {

	var chicken = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    60,
		"april":    20,
	}

	for key, val := range chicken {
		fmt.Println(key, "\t ", val)
	}

	// delete data map
	fmt.Println(len(chicken))
	delete(chicken, "januari")
	fmt.Println(len(chicken))
	fmt.Println(chicken)

	var bulan = map[string]string{
		"januari": "ini adalah bulan januari",
		"april":   "ini bulan april",
	}

	var value, cariBulan = bulan["april"]
	if cariBulan {
		fmt.Println(value)

	} else {
		fmt.Println("bulan tidak ditemukan")
	}
}
