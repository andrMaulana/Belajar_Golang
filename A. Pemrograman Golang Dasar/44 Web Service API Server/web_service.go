package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	ID    string
	Name  string
	Grade int
}

var data = []Student{
	{"S001", "Andri Maulana", 21},
	{"S002", "Ekho Kurniawan", 22},
	{"S003", "Anggi Saputra", 21},
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		student, err := json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(student)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		id := r.FormValue("id")
		var result []byte
		var err error

		for _, each := range data {
			if each.ID == id {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}
		}
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
func main() {
	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	fmt.Println("server running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
