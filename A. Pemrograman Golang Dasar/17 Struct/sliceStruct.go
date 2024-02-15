package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age int
	Grade float64
}
func main() {
	var allStudent = []Student{
		{Name: "Andri Maulana",Age: 23, Grade: 3.44},
		{Name: "Nathan Firdaus", Age: 21, Grade: 3.22},
		{Name: "Josuha", Age: 25, Grade: 3.77},
	}

	for _, student := range allStudent{
		fmt.Println(student.Name, "is age", student.Age, "for grade =>", student.Grade)
	}

	// Anonymouse Slice
	var allPeople = []struct {
		Student
		Address string
	}{
			{Student : Student{"Monard", 24, 3.33}, Address: "kampung cilentah"},
	}

	for _, people := range allPeople{
		fmt.Println(people.Name, "berumur", people.Age)
	}
}
