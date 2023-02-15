package main

/*
Equivalent to:
import "fmt"
import "math"
*/
import (
	"fmt"
	"math"
)

func main() {
	// %g & %e for large exponents, %f otherwise.
	fmt.Printf("Square root of seven is %g.\n", math.Sqrt(7))
}
