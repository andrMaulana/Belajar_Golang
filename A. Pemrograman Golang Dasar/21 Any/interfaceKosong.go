package main

import "fmt"

func main() {
	var secret interface{}

	secret = "roger hunt"
	fmt.Println(secret)

	secret = 23
	fmt.Println(secret)

	secret = true
	fmt.Println(secret)

	secret = []string{"hallo", "test", "dynamic"}
	fmt.Println(secret)

	var data = map[string]interface{}{
		"Name" : "Andri",
		"Age" : 24,
		"Hobbies" : []string{"Football", "Coding", "Design"},
	}

	for k, value := range data{
		fmt.Println(k, " =>", value)
	}
}
