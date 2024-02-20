package main

import "fmt"

func main() {
	var data = "hallo"
	var cast = []byte(data)

	fmt.Printf("%d %d %d %d %d\n", cast[0], cast[1], cast[2], cast[3], cast[4])

	var bite = []byte{104, 97, 108, 108, 111}
	var conv = string(bite)
	fmt.Println(conv)
}
