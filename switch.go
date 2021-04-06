package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		switch i {
		case 1:
			fmt.Println("satu")
		case 2:
			fmt.Println("dua")
		case 3:
			fmt.Println("tiga")
		case 4:
			fmt.Println("empat")
		case 5:
			fmt.Println("lima")
		default:
			fmt.Println("gataw")
		}
	}
}
