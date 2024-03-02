package main

import (
	"fmt"
	"io"
	"os"
)

var path = "/home/andrimlna/Documents/Course/Ebook Course/Nopal Course Golang/A. Pemrograman Golang Dasar/41 File/temp/text.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

// fungsi membuat file baru
func createFile() {
	// deteksi apakah file sudah ada atau belum
	_, err := os.Stat(path)

	// jika belum ada maka akan membuat file baru
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("==> file berhasil dibuat di :", path)
}

// membuat fungsi untuk mengisi file text
func writeFile() {
	// membuka file dan menentuka level akses
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// mengisi file
	_, err = file.WriteString("Hallo Andri Maulana\n")
	if isError(err) {
		return
	}
	_, err = file.WriteString("Selamat belajar python\n")
	if isError(err) {
		return
	}

	// save file
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("==> file berhasil di isi")
}

// membaca isi file
func readFile() {
	// membuka file dan menentukan level akses
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// membaca file
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

	fmt.Println("==> file berhasil dibaca :")
	fmt.Println(string(text))
}

// menghapus file
func deleteFile() {
	err := os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("==> file berhasil di hapus")
}
func main() {
	// createFile()
	// writeFile()
	// readFile()
	deleteFile()
}
