package main

import "fmt"

func main() {
	// penggunaan for biasa
	var fruits = [4]string{"banana", "apple", "melon", "strawberry"}
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("element : %d => %s\n", i, fruits[i])
	}
	// penggunaan menggunakan for - range
	for _, v := range fruits {
		fmt.Printf("isi element : %s\n", v)
	}

	var fruits2 = make([]string, 2)
	fruits2[0] = "salak"
	fruits2[1] = "semangka"
	fmt.Println(fruits2)
}
