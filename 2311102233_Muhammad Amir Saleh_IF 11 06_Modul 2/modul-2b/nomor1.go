package main

import (
	"fmt"
	"strings"
)

func main() {
	amir233 := []string{"merah", "kuning", "hijau", "ungu"}

	numExperiments := 5

	for {
		success := true

		for i := 1; i <= numExperiments; i++ {
			var colors [4]string
			fmt.Printf("Percobaan %d: ", i)
			for j := 0; j < 4; j++ {
				fmt.Scan(&colors[j])
			}

			for j := 0; j < 4; j++ {
				if strings.ToLower(colors[j]) != amir233[j] {
					success = false
					break
				}
			}

			fmt.Printf("Percobaan %d: %s %s %s %s\n", i, colors[0], colors[1], colors[2], colors[3])
		}

		if success {
			fmt.Println("BERHASIL: true")
		} else {
			fmt.Println("BERHASIL: false")
		}

		var again string
		fmt.Print("Lakukan percobaan lagi? (y/n): ")
		fmt.Scan(&again)
		if strings.ToLower(again) != "y" {
			break
		}
		fmt.Println() // Add a blank line for readability
	}
}
