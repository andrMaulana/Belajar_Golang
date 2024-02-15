package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	var messages = make(chan string)

	sayHelloTo := func(who string) {
		data := fmt.Sprintf("Hello %s", who)
		messages <- data
	}
	go sayHelloTo("Andri")
	go sayHelloTo("Danni")
	go sayHelloTo("Irfan")

	message1 := <-messages
	fmt.Println(message1)

	message2 := <-messages
	fmt.Println(message2)

	message3 := <-messages
	fmt.Println(message3)
}
