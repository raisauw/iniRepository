package main

import (
	"fmt"
	"math"
)

func main() {
	var kantong1, kantong2 float64
	amir233 := 0.0

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		_, err := fmt.Scanf("%f %f", &kantong1, &kantong2)
		if err != nil {
			fmt.Println("Input tidak valid. Mohon masukkan dua angka.")
			continue
		}

		if kantong1 < 0 || kantong2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		amir233 = kantong1 + kantong2
		selisih := math.Abs(kantong1 - kantong2)
		akanOleng := selisih >= 9

		fmt.Printf("Sepeda motor pak Andi akan oleng: %t\n", akanOleng)

		if amir233 > 150 {
			fmt.Println("Proses selesai.")
			break
		}
	}
}
