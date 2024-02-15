package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 4)
	<-timer.C
	fmt.Println("finish")
}
