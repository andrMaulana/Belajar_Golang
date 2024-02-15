package main

import "fmt"

func main() {
	data := []string{"Superman", "Batman", "Spiderman", "Ashiapman"}

	for _, each := range data {
		func() {

			// block recover IIEF
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Terjadi Panic pada data", each, "|message :", r)
				} else {
					fmt.Println("Aplikasi berjalan sempurna")
				}
			}()

			panic("Some error")
		}()
	}
}
