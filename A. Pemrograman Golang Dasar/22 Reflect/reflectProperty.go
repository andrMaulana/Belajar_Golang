package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name    string
	Address string
	Age     int
	IsJob   bool
}

func (person *Person) GetPropertiName() {
	var reflectVal = reflect.ValueOf(person)

	if reflectVal.Kind() == reflect.Ptr {
		reflectVal = reflectVal.Elem()
	}

	for i := 0; i < reflectVal.NumField(); i++ {
		fmt.Println("Nama Properti :", reflectVal.Type().Field(i).Name)
		fmt.Println("Tipe Properti :", reflectVal.Type().Field(i).Type)
		fmt.Println("Value Properti :", reflectVal.Field(i).Interface())
		fmt.Println(" ")
	}
}
func (person *Person) SetName(name string) {
	person.Name = name
}
func main() {
	person1 := &Person{
		Name:    "Andri Maulana",
		Address: "dusun cilentah rt 05 rw 02 desa curug",
		Age:     23,
		IsJob:   false,
	}
	person1.GetPropertiName()

	fmt.Println("name :", person1.Name)

	reflectVal := reflect.ValueOf(person1)
	method := reflectVal.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("Martines"),
	})
	fmt.Println(person1.Name)
}
