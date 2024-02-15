package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.Now()
	fmt.Printf("time1 %v\n", time1)

	time2 := time.Date(2011, 12, 24, 10, 20, 0, 0, time.Local)
	fmt.Printf("time2 %v\n", time2)

	// method time.time

	now := time.Now()
	fmt.Println("year :", now.Year(), "Month :", now.Month())
}
