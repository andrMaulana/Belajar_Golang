package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := "123"

	konversi, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println(err.Error())

	}
	fmt.Println(konversi)
}
