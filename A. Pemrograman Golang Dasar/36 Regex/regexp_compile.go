package main

import (
	"fmt"
	"regexp"
)

func main() {

	var text = "banana sugar soup"
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}
	res1 := regex.FindAllString(text, 2)
	fmt.Printf("%#v \n", res1)

	res2 := regex.FindAllString(text, -1)

	fmt.Printf("%#v \n", res2)
	// regex method matchString()

	var isMatch = regex.MatchString(text)
	fmt.Printf("%#v \n", isMatch)

	// Method findString()
	str := regex.FindString(text)
	fmt.Printf("%#v \n", str)

	// method findStringIndex()
	txt := regex.FindStringIndex(text)

	fmt.Printf("%#v \n", txt)

	fmt.Println(text[0:6])

	// method replaceAllString()
	replace := regex.ReplaceAllString(text, "done")
	fmt.Printf("%#v \n", replace)

	// method replaceAllstringfunc()
	replace2 := regex.ReplaceAllStringFunc(text, func(s string) string {
		if s == "banana" {
			return "melon"
		}

		return s
	})

	fmt.Printf("%#v \n", replace2)

	regex2, _ := regexp.Compile(`[a]+`)
	datas := "muhammad nurdin"
	split1 := regex2.Split(datas, -1)
	fmt.Printf("%#v \n", split1)

}
