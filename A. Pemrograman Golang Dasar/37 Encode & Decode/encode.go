package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	var data = "Andri Maulana"

	encodeString := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("hasil encoding :", encodeString)

	decodeToString, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		fmt.Println(err.Error())

	}
	decodeString := string(decodeToString)
	fmt.Println("Hasil encode :", decodeString)

	fmt.Println(strings.Repeat("#", 50))

	datas := "Andri Maulana"
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(datas)))
	base64.StdEncoding.Encode(encoded, []byte(datas))
	fmt.Println("Encoded :", string(encoded))

	var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	_, err = base64.StdEncoding.Decode(decoded, encoded)
	if err != nil {
		fmt.Println(err.Error())
	}
	decodedString := string(decoded)
	fmt.Println("decoded :", decodedString)
	fmt.Println(strings.Repeat("#", 50))
	// url encoding
	var url = "https://google.com/"
	var encodeUrl = base64.URLEncoding.EncodeToString([]byte(url))
	fmt.Println("Url encode :", encodeUrl)

	var decodeUrl, _ = base64.URLEncoding.DecodeString(encodeUrl)
	fmt.Println(string(decodeUrl))
}
