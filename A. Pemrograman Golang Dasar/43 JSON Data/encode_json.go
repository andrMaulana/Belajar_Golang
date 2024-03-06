package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	jsonData := []User{{"Andri", 32}, {"ekho", 32}}

	konvJson, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	toString := string(konvJson)
	fmt.Println(toString)

	fmt.Println(strings.Repeat("*", 50))

	var data []User
	json.Unmarshal([]byte(toString), &data)
	fmt.Println("User 1 :", data[0].FullName)
	fmt.Println("User 2 :", data[1].FullName)
}
