/*
A.1.3. Deklarasi Multi Variabel
*/
package main

import "fmt"

func main() {
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	fmt.Printf("%s %s %s\n", first, second, third)

	var fourth, fifth, sixth string = "empat", "lima", "enam"

	fmt.Printf("%s %s %s\n", fourth, fifth, sixth)
	seventh, eighth, ninth := "tujuh", "delapan", "sembilan"
	fmt.Printf("%s %s %s\n", seventh, eighth, ninth)

	one, isFriday, twoPointTwo, say := 1, false, 2.2, "hallo"
	fmt.Printf("%d %v %.2f %s\n", one, isFriday, twoPointTwo, say)

}
