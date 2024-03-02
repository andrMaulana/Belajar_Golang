package main

import (
	"encoding/json"
	"net/http"
)

var baseUrl = "http://localhost:8080"

type Student struct {
	ID    string
	Name  string
	Grade int
}

func fecthUser() ([]Student, error) {
	var err error
	var client = &http.Client{}
	var data []Student

	request, err := http.NewRequest("GET", baseUrl+"/users", nil)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)

}

func main() {

}
