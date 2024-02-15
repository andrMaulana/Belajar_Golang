package main

import "fmt"

func main() {
	const (
		name      = "andri"
		age       = 24
		isMarried = false
		ipk       = 3.44
	)
	fmt.Printf("hai my name is %s and am %d iam not marrid (%v) am collage in value %.2f\n", name, age, isMarried, ipk)
	const (
		a = "contanta"
		b
	)
	fmt.Println(a)
	fmt.Println(b)
	const satu, dua = 1, 2
	const three, four string = "tiga", "empat"
	fmt.Println(satu, dua)
	fmt.Println(three, four)

}
