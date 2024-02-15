package main

import (
	"fmt"
	"reflect"
)

func main() {
	var data interface{}
	data = 23

	var reflectVal = reflect.ValueOf(data)
	fmt.Println("Tipe data reflect :", reflectVal.Type())
	fmt.Println("Value Reflect :", reflectVal.Interface())
}
