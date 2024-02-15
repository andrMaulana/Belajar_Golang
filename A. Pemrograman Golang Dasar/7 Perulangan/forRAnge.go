package main

import (
	"fmt"
	"strings"
)

func main() {
	var xs string = "123" // for range untuk data string
	for i, v := range xs {
		fmt.Printf("%d => %d\n", i, v)
	}
	fmt.Println(strings.Repeat("#", 10))
	fmt.Println("data array")
	var ys = [5]int{10, 20, 30, 40, 50} // for range untuk data array
	for i, v := range ys {
		fmt.Printf("%d => %d\n", i, v)
	}

	fmt.Println(strings.Repeat("#", 10))
	fmt.Println("data slice")
	var zx = ys[0:2] // fo range untuk data slice
	for i, v := range zx {
		fmt.Printf("%d => %d\n", i, v)
	}

	fmt.Println(strings.Repeat("#", 10))
	fmt.Println("data map")
	var kvs = map[byte]int{'a': 1, 'b': 2, 'c': 3} // for range untuk data map
	for k, v := range kvs {
		fmt.Printf("%d => %d\n", k, v)
	}

	for range kvs {
		fmt.Println("done")
	}

}
