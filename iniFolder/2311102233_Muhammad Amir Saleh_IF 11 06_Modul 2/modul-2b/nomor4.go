package main

import (
	"fmt"
	"math"
)

func f(k float64) float64 {
	amir233 := math.Pow(4*k+2, 2)
	denominator := (4*k + 1) * (4*k + 3)
	return amir233 / denominator
}

func approximateSqrt2(K int) float64 {
	result := 1.0
	for k := 0; k <= K; k++ {
		result *= f(float64(k))
	}
	return result
}

func main() {
	var K int

	fmt.Print("Nilai K = ")
	fmt.Scan(&K)

	fK := f(float64(K))
	fmt.Printf("Nilai f(K) = %.10f\n", fK)

	sqrt2 := approximateSqrt2(K)
	fmt.Printf("Nilai akar 2 = %.10f\n", sqrt2)
}
