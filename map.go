package main

import "fmt"

func main() {
	mahasiswa := make(map[string]string)

	mahasiswa["1001"] = "Budi"
	mahasiswa["1002"] = "Anto"
	mahasiswa["1003"] = "Cece"

	fmt.Println(mahasiswa)

	for nim, name := range mahasiswa {
		fmt.Println(nim, "bernama", name)
	}
}
