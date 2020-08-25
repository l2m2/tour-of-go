package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := x / 2
	for {
		if diff := z*z - x; diff < 0.00001 && diff > -0.00001 {
			return z
		} else {
			z -= (z*z - x) / (2 * z)
			fmt.Println(z)
		}
	}
}

func main() {
	fmt.Println(Sqrt(80))
}
