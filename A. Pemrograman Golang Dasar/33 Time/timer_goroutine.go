package main

import (
	"fmt"
	"os"
	"time"
)

func Timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func Watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\nwaktu habis tidak ada jawaban dalam waktu ", timeout, "detik")
	os.Exit(0)
}

func main() {
	timeout := 2
	ch := make(chan bool)

	go Timer(timeout, ch)
	go Watcher(timeout, ch)

	var soal1 string
	fmt.Print("hasil dari 725/25 ? = ")
	fmt.Scan(&soal1)

	if soal1 == "29" {
		fmt.Println("Hore Jawaban Benar!!!")

	} else {
		fmt.Println("Yahh Jawaban Salah!")
	}
}
