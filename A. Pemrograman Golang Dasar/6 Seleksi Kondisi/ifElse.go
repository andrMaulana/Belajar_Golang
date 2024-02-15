package main

import "fmt"

func main() {
	var nilai int = 5

	if nilai == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if nilai >= 5 {
		fmt.Println("lulus")
	} else if nilai == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", nilai)
	}

	// Seleksi kondisi dengan variabel temporary
	var point float64 = 8840.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good!\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}
}
