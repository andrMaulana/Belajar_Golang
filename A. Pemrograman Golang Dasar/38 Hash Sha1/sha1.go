package main

import (
	"crypto/sha1"
	"fmt"
	"strings"
	"time"
)

func main() {
	test := "this is secret"
	sha := sha1.New()
	sha.Write([]byte(test))
	encryp := sha.Sum(nil)
	res := fmt.Sprintf("%x \n", encryp)
	fmt.Println("Hasil encrypt", res)

	fmt.Println(strings.Repeat("#", 50))
	var name = "Andri maulana"
	fmt.Printf("original 1 : %s\n\n", name)

	hash1, salt1 := doHashSalting(name)
	fmt.Printf("hashed 1 : %s\n\n", hash1)

	hash2, salt2 := doHashSalting(name)
	fmt.Printf("hashed 2 : %s\n\n", hash2)

	hash3, salt3 := doHashSalting(name)
	fmt.Printf("hashed 3 : %s\n\n", hash3)

	_, _, _ = salt1, salt2, salt3
}

// metode salt
func doHashSalting(text string) (string, string) {
	salt := fmt.Sprintf("%d", time.Now().UnixNano())
	salText := fmt.Sprintf("text : '%s', salt : %s", text, salt)
	fmt.Println(salText)
	sha := sha1.New()
	sha.Write([]byte(salText))
	encrypted := sha.Sum(nil)
	return fmt.Sprintf("%x", encrypted), salt
}
