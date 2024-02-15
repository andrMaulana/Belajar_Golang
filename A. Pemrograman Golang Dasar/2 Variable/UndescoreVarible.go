/*
A.1.4. Variabel Underscore _

	Go memiliki aturan unik yang jarang dimiliki bahasa lain, yaitu tidak boleh ada satupun variabel yang menganggur. Artinya, semua variabel yang dideklarasikan harus digunakan. Jika ada variabel yang tidak digunakan tapi dideklarasikan, error akan muncul pada saat kompilasi dan program tidak akan bisa di-run
*/
package main

import "fmt"

func main() {
	_ = "belajar golang"
	_ = "golang itu mudah"
	name, _ := "andri", "maulana"
	fmt.Printf(name)
}
