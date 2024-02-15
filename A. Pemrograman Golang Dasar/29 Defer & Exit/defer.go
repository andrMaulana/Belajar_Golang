package main

import (
	"fmt"
)

func orderSomeFood(menu string) {
	defer fmt.Println("terima kasih pesananya")

	if menu == "pizza" {
		fmt.Println("pesanan yang tepat")
		fmt.Println(menu, "akan disiapkan!")
		return
	}
	fmt.Println("pesanan anda adalah :", menu)
}

func main() {

	fmt.Println("selamat datang")
	orderSomeFood("hamburger")
	orderSomeFood("pizza")

	// defer block fungsi

	number := 3
	if number == 3 {
		fmt.Println("Hallo 1")
		func() {
			fmt.Println("Hallo 3")
		}()
	}

	fmt.Println("Hallo 2")
}
