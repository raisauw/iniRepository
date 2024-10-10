package main

import (
	"fmt"
)

func amir233(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func main() {
	var celsius float64

	fmt.Print("Masukkan suhu dalam Celsius: ")
	fmt.Scan(&celsius)

	fahrenheit := amir233(celsius)
	fmt.Println(fahrenheit)
}
