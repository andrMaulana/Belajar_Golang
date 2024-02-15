/*
A.1.1. Deklarasi Variabel Beserta Tipe Data

Pada kode di atas bisa dilihat bagaimana sebuah variabel dideklarasikan dan diisi
nilainya. Keyword var
digunakan untuk membuat variabel baru.
Skema penggunaan keyword var:

	var <nama-variabel> <tipe-data>
	var <nama-variabel> <tipe-data> = <nilai>

Contoh:

	var lastName string
	var firstName string = "john"
*/
package main

import "fmt"

func main() {
	var firstName string = "Andri"
	var lastName string
	lastName = "Maulana"
	fmt.Printf("Hallo Andri Maulana!\n")               // should be "Hallo Andri Maulana!"
	fmt.Printf("Hallo %s %s! \n", firstName, lastName) // should be "Hallo Andri Maulana!"
	fmt.Println("Hallo", firstName, lastName+"!")      // should be "Hallo Andri Maulana!"
}
