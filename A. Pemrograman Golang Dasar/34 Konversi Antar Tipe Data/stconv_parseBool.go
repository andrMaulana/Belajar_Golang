package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "true"

	conv, err := strconv.ParseBool(str)

	if err != nil {
		fmt.Println(err.Error())

	}

	fmt.Println(conv)
}
