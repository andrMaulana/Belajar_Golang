package main

import (
	"fmt"
	"runtime"
)

func printMessage(message chan string) {
	fmt.Println(<-message)
}
func main() {
	runtime.GOMAXPROCS(2)

	messages := make(chan string)

	for _, each := range []string{"John", "Arney", "gilbert"} {
		go func(who string) {
			var datas = fmt.Sprintf("Hallo %s", who)
			messages <- datas
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages)
	}
}
