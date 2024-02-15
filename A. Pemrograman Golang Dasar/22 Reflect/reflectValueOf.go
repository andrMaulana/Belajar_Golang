package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	var data interface{}
	data = "Andri"
	var reflectValString = reflect.ValueOf(data)

	if reflectValString.Kind() == reflect.String {
		fmt.Println("Tipe data Reflect :", reflectValString.Type())
		fmt.Println("Data String", reflectValString.String())
	} else {
		fmt.Println("Data gagal diambil karena bukan tipe data string")
	}
	fmt.Println(strings.Repeat("#", 50))
	data = 242
	reflectValInt := reflect.ValueOf(data)

	if reflectValInt.Kind() == reflect.Int {
		fmt.Println("Tipe data reflect :", reflectValInt.Type())
		fmt.Println("Data Int :", reflectValInt.Int())
	} else {
		fmt.Println("Gagal")
	}
	fmt.Println(strings.Repeat("#", 50))
	data = true
	reflectValBool := reflect.ValueOf(data)

	if reflectValBool.Kind() == reflect.Bool {
		fmt.Println("Tipe data reflect :", reflectValBool.Type())
		fmt.Println("Data Bool :", reflectValBool.Bool())
	} else {
		fmt.Println("Gagal")
	}
}
