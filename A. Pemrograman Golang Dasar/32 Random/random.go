package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(length int) string {
	b := make([]rune, length)

	for i := range b {
		b[i] = letter[randomizer.Intn(len(letter))]
	}

	return string(b)
}

func main() {

	fmt.Println(strings.Repeat("#", 10))
	fmt.Println("ke -1 :", randomizer.Int())
	fmt.Println("ke -2 :", randomizer.Int())
	fmt.Println("ke -3 :", randomizer.Int())

	fmt.Println(strings.Repeat("#", 10))
	fmt.Println("random data int => ", randomizer.Int())
	fmt.Println("random data float => ", randomizer.Float32())
	fmt.Println("random data uint => ", randomizer.Uint32())

	// angka random dengan jarak tertentu
	fmt.Println(strings.Repeat("#", 10))
	fmt.Println("random data jarak => ", randomizer.Intn(100))

	// random string
	fmt.Println(strings.Repeat("#", 10))
	fmt.Println("random data string => ", randomString(50))
}
