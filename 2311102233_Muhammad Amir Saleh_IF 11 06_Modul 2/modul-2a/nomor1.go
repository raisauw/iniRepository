package main

import (
	"fmt"
)

func main() {
	var (
		satu, dua, amir233 string
		temp               string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&amir233)
	fmt.Println("Output awal = " + satu + " " + dua + " " + amir233)
	temp = satu
	satu = dua
	dua = amir233
	amir233 = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + amir233)
}
