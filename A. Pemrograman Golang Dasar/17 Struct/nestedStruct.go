package main

import "fmt"
type Student struct{
	person struct{
		Name string
		Age string
	}
	Grade int
	Hobbies []string
}
func main() {
	var student = Student{
		struct{Name string; Age string}{"Andri Maulana", "24 Tahun"},
		Grade : 32,
		Hobbies: ["coding", "catur", "football"],
	}
	fmt.Println(student)
}

