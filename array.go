package main

import "fmt"

func main() {
	var kata [3]string

	kata[0] = "Appa"
	kata[1] = "Yip"
	kata[2] = "Yip"

	for i := 0; i < len(kata); i++ {
		fmt.Println(kata[i])
	}

	for index, value := range kata {
		fmt.Println("Index", index, "=", value)
	}
}
