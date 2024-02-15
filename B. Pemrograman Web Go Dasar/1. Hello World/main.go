package main

import (
	"fmt"
	"net/http"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	message := "Welcome"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	message := "Hello World"
	w.Write([]byte(message))
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := "Welcome! Selamat belajar golang web"
		w.Write([]byte(message))
	})
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	address := "localhost:9000"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
