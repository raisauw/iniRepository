package main

import "fmt"

func volumeBola(jari float64) float64 {
	const pi = 3.1415926535
	return (4.0 / 3.0) * pi * jari * jari * jari
}

func luasPermukaanBola(jari float64) float64 {
	const pi = 3.1415926535
	return 4 * pi * jari * jari
}

func main() {
	var amir233 float64

	fmt.Print("Masukkan amir233: ")
	fmt.Scan(&amir233)

	volume := volumeBola(amir233)
	surfaceArea := luasPermukaanBola(amir233)

	fmt.Printf("Bola dengan jari-jari %.2f memiliki volume %.4f dan luas kulit %.4f\n", amir233, volume, surfaceArea)
}
