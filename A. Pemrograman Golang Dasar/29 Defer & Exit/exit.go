package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Hallo")
	os.Exit(1)
	fmt.Println("Selamat Datang")
}
