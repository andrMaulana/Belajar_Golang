package main

import "fmt"

type Student struct{
	Name string
	Grade float64
}
func main() {
	var s1 = struct{
	Student
	Age int
	}{}

	s1.Name = "Andri Maulana"
	s1.Grade = 3.34
	s1.Age = 23

	fmt.Println("Nama :", s1.Name)
	fmt.Println("Grade :", s1.Grade)
	fmt.Println("Age :", s1.Age)
}
