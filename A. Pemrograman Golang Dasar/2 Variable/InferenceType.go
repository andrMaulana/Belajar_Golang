/*
A.1.2 Deklarasi Variable Tanpa Tipe Data

	Selain mengadopsi maifest typing, go juga mengadopsi type inference, yaitu metode deklarasi variable yang datanya ditentukan oleh tipe data nilainya, cara kontradiktif dengan cara pertama.Dengan metode jenis ini, keyword var dan tipe data tidak perlu ditulis.

	Tanda :=
	hanya digunakan sekali di awal pada saat deklarasi. Untuk assignment nilai selanjutnya harus menggunakan tanda
	= , contoh:
	lastName := "wick"
	lastName = "ethan"
	lastName = "bourne"
	Perlu diketahui, deklarasi menggunakan := hanya bisa dilakukan di dalam blok fungsi.
*/
package main

import "fmt"

func main() {
	// menggunakan keyword var tanpa menuliskan tipe data
	var firstName = "Andri"
	// tanpa menggunakan keyword var diganti dengan menggunakan :=
	lastName := "Maulana"

	fmt.Printf("Hallo %s %s!\n", firstName, lastName)

}
