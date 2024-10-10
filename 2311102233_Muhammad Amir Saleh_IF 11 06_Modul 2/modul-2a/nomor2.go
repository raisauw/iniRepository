package main

import (
	"fmt"
)

func main() {
	var amir233 int
	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&amir233)

	if (amir233%4 == 0 && amir233%100 != 0) || (amir233%400 == 0) {
		fmt.Printf("true")
	} else {
		fmt.Printf("false")
	}
}
