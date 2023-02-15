package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	var prev float64
	for i := 1; i < 10; i += 1 {
		z -= (z*z - x) / (2 * z)
		if z == prev {
			break
		}
		fmt.Println("iteration", i, "->", z)
		prev = z
	}
	return z
}

func main() {
	fmt.Println("result:", Sqrt(4))
}
