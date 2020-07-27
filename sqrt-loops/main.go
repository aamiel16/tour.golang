package main

import (
	"fmt"
	"math"
)

type SqrtError struct {
	num float64
}

func (e SqrtError) Error() string {
	return fmt.Sprintf("Cannot call Sqrt on negative number = %f", e.num)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, SqrtError{x}
	}

	const MAX_DELTA = 0.00000001

	var (
		diff float64 = math.MaxFloat64
		z float64 = 1.0
		prev float64
	)

	for diff > MAX_DELTA {
		prev = z
		z -= (z*z - x) / (2*z) // Newton-Raphsons method for approximation
		diff = math.Abs(prev - z)
	}

	return z, nil;
}


func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(-2))
}
