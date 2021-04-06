package main

import "fmt"

func main() {
	slice1 := make([]string, 3)
	slice1[0] = "joko"
	slice1[1] = "owi"
	slice1[2] = "dodo"

	slice2 := append(slice1, "budi", "amir")
	slice2[0] = "mama"

	fmt.Println("======================")
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println("======================")

	slice3 := make([]string, 3)
	copy(slice3, slice1)

	slice1[0] = "bambang"

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println("======================")

	slice4 := slice1[:]
	slice4[2] = "dudu"

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println("======================")
}
