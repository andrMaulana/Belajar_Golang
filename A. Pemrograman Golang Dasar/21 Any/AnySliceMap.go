package main

import "fmt"

func main() {
	data := []map[string]any{
		{"Name" : "Andri", "Age" : 23},
		{"Name" : "Garry", "Age" : 22},
		{"Name" : "John", "Age" : 21},
	}

	for _, val := range data{
		fmt.Println(val["Name"], "is age", val["Age"])
	}

	fruits := []interface{}{
		map[string]interface{}{
			"Name" : "Strawberry",
			"Total" : 100,
		},
		[]string{"Manggo", "Apple", "Pineaple", "Grape"},
		"Watermelon",
	}

	for _, val := range fruits{
		fmt.Println(val)
	}
}
