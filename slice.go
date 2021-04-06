package main

import "fmt"

func main() {
	names := [5]string{
		"adit",
		"sopo",
		"jarwo",
		"denis",
		"budi",
	}

	fmt.Println(names)

	slice1 := names[:3]
	fmt.Println(slice1)

	slice1[0] = "malih"

	fmt.Println(names)
	fmt.Println(slice1)
}
