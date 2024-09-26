package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var path = "/home/andrimlna/Documents/Course/Ebook Course/Nopal Course Golang/A. Pemrograman Golang Dasar/41 File/temp/test.json"

type ListData struct {
	ID           int    `json:"id"`
	PropertyName string `json:"property_name"`
	Url          string `json:"url"`
}

func readJson() []ListData {
	var list_data []ListData
	var respData []ListData

	fileJson, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer fileJson.Close()
	byteValue, _ := ioutil.ReadAll(fileJson)
	json.Unmarshal(byteValue, &list_data)
	for i := 0; i < len(list_data); i++ {
		lst := ListData{
			ID:           list_data[i].ID,
			PropertyName: list_data[i].PropertyName,
			Url:          list_data[i].Url,
		}

		respData = append(respData, lst)
	}
	return respData
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func CreateFile() {
	// cek lokasi file
	_, err := os.Stat(path)

	// create file
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("==> file successfully created", path)
}

func writeFile() {
	// open file
	file, err := os.OpenFile(path, os.O_RDWR, 0664)
	if isError(err) {
		return
	}
	defer file.Close()

	// isi file
	_, err = file.WriteString("How to be a backend engineering\n")
	if isError(err) {
		return
	}
	_, err = file.WriteString("The backend engineering to be hard to keep leaning\n")
	if isError(err) {
		return
	}
	// sync for save
	err = file.Sync()
	fmt.Println("==>> file successfully synced", path)
}
func readFile() {
	// read file
	file, err := os.OpenFile(path, os.O_RDONLY, 0664)
	if isError(err) {
		return
	}
	defer file.Close()
	// read file
	text := make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if isError(err) {
				break
			}
		}
		if n == 0 {
			break
		}
	}
	if isError(err) {
		return
	}
	fmt.Println("file successfully read")
	fmt.Println(string(text))

}

func deleteFile() {
	err := os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("===>> file successfully deleted")
}

func main() {
	// CreateFile()
	// writeFile()
	// readFile()
	// deleteFile()

	fmt.Println(readJson())
}
