package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z = 1.0
	for math.Abs(x-z*z) > 1e-8 {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
