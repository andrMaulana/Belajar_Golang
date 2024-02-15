package main

import "fmt"

type Person struct{
	Name string
	Age int
}
func main() {
	var person any = &Person{Name: "Andri Maulana", Age: 23}
	fmt.Println(person.(*Person).Name)
}
