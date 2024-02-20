package main

import (
	"flag"
	"fmt"
)

func main() {
	// var name = flag.String("Name", "null", "masukan nama")
	// var age = flag.Int64("Umur", 0, "masukan umur")

	// flag.Parse()
	// fmt.Printf("Nama : %v\n", *name)
	// fmt.Printf("Umur : %v\n", *age)

	// metode 2
	var data1 = flag.String("name", "anonymouse", "type your name")
	var data2 string
	flag.StringVar(&data2, "gender", "male", "type your gender")

	flag.Parse()
	fmt.Println(*data1)
	fmt.Println(data2)
}
