package main

import (
	"fmt"
	"strings"
)

func main() {
	var amir233 string
	var jumlahBunga int

	for {
		var namaBunga string
		jumlahBunga++
		fmt.Printf("Bunga %d: ", jumlahBunga)
		fmt.Scan(&namaBunga)

		if strings.ToUpper(namaBunga) == "SELESAI" {
			jumlahBunga--
			break
		}

		if amir233 == "" {
			amir233 = namaBunga
		} else {
			amir233 += " - " + namaBunga
		}
	}

	fmt.Println("amir233:", amir233)
	fmt.Println("Bunga:", jumlahBunga)
}
