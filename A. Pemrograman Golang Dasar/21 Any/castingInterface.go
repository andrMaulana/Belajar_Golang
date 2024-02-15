package main

import (
	"fmt"
	"strings"
)

func main() {
	var secret interface{} = 12
	var multiple = secret.(int) * 12
	fmt.Println(multiple)

	var datas any = []string{"John", "Gary", "dump"}
	joins := strings.Join(datas.([]string), ", ")
	fmt.Println(joins)

	dataMap := map[string]any{"name" : "garry","Age": 24}
	for _,val := range dataMap{
		fmt.Printf("Hallo %s your Age is %d", val.(string), val.(int))
	}
}
