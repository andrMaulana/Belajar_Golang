package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlParsing := "http://kalipare.com:80/hello?name=john_wick&age=27"
	u, err := url.Parse(urlParsing)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("data url :", urlParsing)
	fmt.Printf("data pasing : %s://%s/%s?name=%s&age=%s\n", u.Scheme, u.Host, u.Path, u.Query()["name"][0], u.Query()["age"][0])
}
