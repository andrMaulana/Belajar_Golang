package main

import "fmt"


type Student struct{
	Name string
	Grade int
}
func main() {
	var s1 Student

	s1.Name = "Andri Maulana"
	s1.Grade = 22

	var s2 = Student{"Agung Herman", 21}

	var s3 = Student{Name: "Anggita"}

	fmt.Println("Nama :", s1.Name)
	fmt.Println("Grade :", s1.Grade)
	fmt.Println("Nama :", s2.Name)
	fmt.Println("Nama : ", s3.Name)
}
