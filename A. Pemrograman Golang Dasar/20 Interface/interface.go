package main

import (
	"fmt"
	"math"
	"strings"
)

type hitung interface{
	luas() float64
	keliling() float64
}

type Lingkaran struct{
	Diameter float64
}

func (l Lingkaran) jarijari() float64  {
	return l.Diameter / 2
}

func (l Lingkaran) luas() float64  {
	return math.Pi * math.Pow(l.jarijari(), 2)
}

func (l Lingkaran) keliling() float64{
	return math.Pi * l.Diameter
}

type Persegi struct{
	Sisi float64
}

func (p Persegi) luas() float64{
	return math.Pow(p.Sisi, 2)
}

func (p Persegi) keliling() float64{
	return p.Sisi * 4
}
func main() {
	var bangunDatar hitung

	bangunDatar = Lingkaran{14.0}
	fmt.Println(strings.Repeat("=",5), "Lingkaran", strings.Repeat("=", 5))
	fmt.Println("Luas Lingkaran =", bangunDatar.luas())
	fmt.Println("Keliling Lingkaran =", bangunDatar.keliling())
	fmt.Println("Jari Jari Lingkaran =", bangunDatar.(Lingkaran).jarijari())

	bangunDatar = Persegi{10.0}
	fmt.Println(strings.Repeat("=", 5), "Persegi", strings.Repeat("=", 5))
	fmt.Println("Luas Persegi =", bangunDatar.luas())
	fmt.Println("Keliling Persegi =", bangunDatar.keliling())
}
