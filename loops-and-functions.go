package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {
	z = 2.0
	y := float64(0)

	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)

		if math.Abs(z-y) < 1e-10 {
			break
		}

		y = z
	}

	return
}

func main() {
	fmt.Println(Sqrt(0.85))
	fmt.Println(math.Sqrt(0.85))
}
