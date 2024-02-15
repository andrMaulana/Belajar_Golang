package main

import (
	"fmt"
	"strings"
)

type Student struct{
	Name string
	Age int
}

func (s Student) changeName1(name string) {
	fmt.Println("---->on change Name 1")
	s.Name = name
}

func (s *Student) changeName2(name string){
	fmt.Println("---->on change Name 2")
	s.Name = name
}

func (s Student) SayHello(){
	fmt.Println("Hallo", s.Name)
}

func (s Student) GetName(i int) string{
	return strings.Split(s.Name, " ")[i-1]
}
func main() {
	var s1 = Student{Name: "Andri Maulana", Age: 23}

	s1.SayHello()
	fmt.Println(s1.GetName(2))
	s1.changeName1("Doni Irawan")
	fmt.Println(s1.Name)
	s1.changeName2("Doni Irawan")
	fmt.Println(s1.Name)

	s2 := &Student{Name: "Ethan",Age: 23}
	fmt.Println(s2)
}
