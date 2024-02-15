package main

import (
	"fmt"
	"strings"
)

func main() {

	hobbies := []string{"Coding", "Football", "Gaming"}
	yourHobbies("Andri", hobbies...)

}

func yourHobbies(names string, hobbies ...string) {
	var yourHobbie = strings.Join(hobbies, ", ")

	fmt.Printf("Hallo My Name Is %s\n", names)
	fmt.Printf("My Hobbies Is %s\n", yourHobbie)
}
