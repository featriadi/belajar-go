package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int32 = 10
	var y int64 = int64(x)

	fmt.Println(y)

	var z float64 = 3.9
	fmt.Println(z)

	var a int32 = int32(z)
	fmt.Println(a)

	var nilai string = "100"
	nilaiInt, _ := strconv.Atoi(nilai)

	fmt.Println(nilaiInt)

	nilaiString := strconv.Itoa(-100)
	fmt.Println(nilaiString)
}
