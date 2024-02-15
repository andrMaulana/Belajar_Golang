/*
A.10.1. Tipe Data Numerik Non-Desimal

	Tipe data numerik non-desimal atau non floating point di Go ada beberapa jenis. Secara umum ada 2 tipe data kategori ini yang perlu diketahui :
	uint , tipe data untuk bilangan cacah (bilangan positif).
	int , tipe data untuk bilangan bulat (bilangan negatif dan positif).
*/
package main

import "fmt"

func main() {
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("positive Number = %d\n", positiveNumber)
	fmt.Printf("Negative Number = %d\n", negativeNumber)
}
