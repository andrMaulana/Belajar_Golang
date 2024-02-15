package main

import "fmt"

func main() {
	var chikens = []map[string]string{
		{
			"name":  "fried chicken",
			"harga": "10.000",
		},
		{
			"name":  "chicken nuggets",
			"harga": "20.000",
		},
		{
			"name":  "maybe ayam",
			"harga": "12.000",
		},
	}

	for _, chicken := range chikens {
		fmt.Println(chicken["name"], chicken["harga"])
	}
}
