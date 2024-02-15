package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)

	time.AfterFunc(4*time.Second, func() {
		<-time.After(time.Second * 2)
		fmt.Println("exfired")
		ch <- true
	})

	fmt.Println("start")
	<-ch
	fmt.Println("finish")
	fmt.Printf("%v\n", ch)
}
