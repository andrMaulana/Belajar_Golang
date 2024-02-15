package main

import (
	"fmt"
)

type Person struct{
	Name string
	Age int
}

type Student struct{
	Grade int
	Person
}
func main() {
	var s1 = Student{}
	s1.Name = "Andri Maulana"
	s1.Age = 22
	s1.Grade = 12

	fmt.Println("Nama :", s1.Name)
	fmt.Println("Age :", s1.Age)
	fmt.Println("Grade :", s1.Grade)
}
