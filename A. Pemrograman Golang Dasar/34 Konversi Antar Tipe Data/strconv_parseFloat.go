package main

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "12.21"

	numFloat, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(numFloat)
}
