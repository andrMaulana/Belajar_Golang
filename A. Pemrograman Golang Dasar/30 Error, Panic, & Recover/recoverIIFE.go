package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi Error:", r)
		} else {
			fmt.Println("apps berjalan sempurna")
		}
	}()

	panic("aplikasi error tidak dapat berjalan")
}
