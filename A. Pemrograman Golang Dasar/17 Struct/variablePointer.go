package main

import "fmt"


type Student struct{
	Name string
	Grade int
}
func main() {
	var s1 = Student{Name: "Andri Maulana",
	Grade: 22,}

	var s2 *Student = &s1

	fmt.Println("Nama Student s1 =", s1.Name)
	fmt.Println("Nama Studen s2 =", s2.Name)

	s2.Name = "Jaka"
	fmt.Println("Nama Student s1 =", s1.Name)
	fmt.Println("Nama Student s2 =",s2.Name)

}
