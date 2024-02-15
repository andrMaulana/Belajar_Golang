package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := "1234"
	conv, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(conv)

	// konveri ke basis biner
	number = "1010"

	conv, err = strconv.ParseInt(number, 2, 8)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(conv)
}
