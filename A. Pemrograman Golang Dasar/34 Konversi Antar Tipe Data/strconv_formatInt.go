package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := int64(24)

	str := strconv.FormatInt(num, 8)
	fmt.Println(str)
}
