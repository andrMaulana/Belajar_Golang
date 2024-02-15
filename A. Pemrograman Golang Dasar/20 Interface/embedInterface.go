package main

import (
	"fmt"
	"math"
	"strings"
)

type hitung2d interface{
	Luas() float64
	Keliling() float64
}

type hitung3d interface{
	Volume() float64
}

type hitung interface{
	hitung2d
	hitung3d
}

type Kubus struct{
	sisi float64
}

func (k *Kubus) Luas() float64{
	return math.Pow(k.sisi, 3)
}

func (k *Kubus) Keliling() float64{
	return math.Pow(k.sisi, 2) * 6
}

func (k *Kubus) Volume() float64{
	return k.sisi * 12
}

func main() {
	var bangunruang hitung = &Kubus{4}

	fmt.Println(strings.Repeat("=", 5), "Kalkulasi Kubus", strings.Repeat("=", 5))
	fmt.Println("Luas Kubus =", bangunruang.Luas())
	fmt.Println("Keliling Kubus =", bangunruang.Keliling())
	fmt.Println("Volume Kubus =", bangunruang.Volume())

}
