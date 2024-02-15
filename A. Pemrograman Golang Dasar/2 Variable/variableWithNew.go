/*
A.1.6. Deklarasi Variabel Menggunakan Keyword new

	Keyword new digunakan untuk membuat variabel pointer dengan tipe data tertentu. Nilai data default-nya akan menyesuaikan tipe datanya.
*/
package main

import "fmt"

func main() {
	name := new(string)
	fmt.Println(name)
	fmt.Println(*name)
}
