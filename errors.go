package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	z := float64(2)
	y := float64(0)

	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)

		if math.Abs(z-y) < 1e-10 {
			break
		}

		y = z
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
