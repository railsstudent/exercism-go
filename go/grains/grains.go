package grains

import "errors"
import "math"

func Square(num int) (uint64, error) {
	if num < 1 || num > 64 {
		return 0, errors.New("Number of squares is out of range")
	}
	return uint64(math.Pow(2.0, float64(num-1))), nil
}

// Total calculates number of grains in 64 chess squares
func Total() uint64 {
	// geometric series [1 * (1 - 2^64)] / (1 - 2) = (1 - 2^64) / -1
	// 2^64 - 1
	grains, _ := Square(64)
	return 2*grains - 1
}
