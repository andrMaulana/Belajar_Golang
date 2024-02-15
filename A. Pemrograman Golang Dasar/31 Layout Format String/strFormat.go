package main

import (
	"fmt"
	"strings"
)

type Student struct {
	Name       string
	Height     float32
	Age        int
	IsGraduate bool
	Hobbies    []string
}

func main() {
	var student = Student{
		Name:       "Andri Maulana",
		Height:     180.1,
		Age:        24,
		IsGraduate: false,
		Hobbies:    []string{"Footbal", "Coding"},
	}

	//  Layout Format %b <- digunakan untuk basis biner
	fmt.Printf("notasi format b => %b\n", student.Age)

	// Layout Format %c <- digunakan untuk format data unicode
	fmt.Printf("notasi format c => %c\n", 1400)
	fmt.Printf("notasi format c => %c\n", 1235)

	// Layout format %d <- digunakan untuk format data numerik
	fmt.Printf("notasi format d => %d\n", student.Age)

	// layout format %e | %E <- digunakan untuk format data numerik desimal ke bentuk notasi numerik standar
	fmt.Printf("notasi format e => %e\n", student.Height)
	fmt.Printf("notasi format E => %E\n", student.Height)

	// layout format %f | %F <- untuk format data numerik desimal dengan lebar default 6 digit
	fmt.Printf("notasi format f => %f\n", student.Height)
	fmt.Printf("notasi format f => %.9f\n", student.Height)
	fmt.Printf("notasi format f => %.2f\n", student.Height)
	fmt.Printf("notasi format f => %.f\n", student.Height)

	// layout format %g | %G <- untuk format data desimal yang cukup besar
	fmt.Printf("notasi format e => %e\n", 0.123123123123)
	fmt.Printf("notasi format f => %f\n", 0.123123123123)
	fmt.Printf("notasi format g => %g\n", 0.123123123123)

	// layout format %o <- untuk format data basis oktal

	fmt.Printf("notasi format o => %o\n", student.Age)

	// layout format %p <- untuk format data pointer
	fmt.Printf("notasi format p => %p\n", &student.Name)

	// layout format %q <- untuk escape string
	fmt.Printf("%q\n", `" name \ height "`)

	// layout format %t <- untuk format data boolean
	fmt.Printf("notasi format t => %t\n", student.IsGraduate)

	// layout foramt %T <- untuk mengambil tipe variable

	fmt.Printf("notasi format T => %T\n", student.Name)
	fmt.Printf("notasi format T => %T\n", student.Height)
	fmt.Printf("notasi format T => %T\n", student.Age)
	fmt.Printf("notasi format T => %T\n", student.IsGraduate)
	fmt.Printf("notasi format T => %T\n", student.Hobbies)

	// layout format %v <- untuk melakukan format pada data tipe apa saja
	fmt.Printf("notasi format v => %v\n", student)

	// layout format %+v <- digunakan untuk memformat struct , mengembalikan nilai beserta propertinya
	fmt.Printf("notas format +v => %+v\n", student)

	// layout format %#v <- mengembalikan nilai di setiap properti struct
	fmt.Printf("notasi format #v => %#v\n", student)

	fmt.Println(strings.Repeat("#", 10))
	var data = struct {
		name   string
		height float32
	}{
		name:   "wick",
		height: 172.1,
	}

	fmt.Printf("%#v\n", data)
	fmt.Println(strings.Repeat("#", 10))

	// layout format %x | %X <- digunakan untuk format data numerik ke string numerik basis 16
	fmt.Printf("notasi format x => %x\n", data.height)

	var d = data.name
	fmt.Printf("notasi format X => %X%X%X%X\n", d[0], d[1], d[2], d[3])

	// layout format %% <- untuk string karakter format
	fmt.Printf("notasi format karakter => %%\n")

}
