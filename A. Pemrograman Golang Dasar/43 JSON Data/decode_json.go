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
	jsonString := `{"Name" : "Andri Maulana", "Age" : 23}`

	var data User
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("User :", data.FullName)
	fmt.Println("Age :", data.Age)

	fmt.Println(strings.Repeat("*", 50))
	var data2 interface{}
	json.Unmarshal([]byte(jsonString), &data2)

	dataJson := data2.(map[string]interface{})
	fmt.Println("user", dataJson["Name"])
	fmt.Println("user", dataJson["Age"])

	fmt.Println(strings.Repeat("*", 50))

	arrayJson := `[
		{"Name" : "Andri", "Age" : 22},
		{"Name" : "Khanedi", "Age" : 32}
	]`

	var data3 []User

	err = json.Unmarshal([]byte(arrayJson), &data3)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("user :", data3[0].FullName)
	fmt.Println("user :", data3[1].FullName)
}
