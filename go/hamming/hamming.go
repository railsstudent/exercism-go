package hamming

import "errors"

// Distance calculates hamming distance
func Distance(a, b string) (int, error) {
	lenA, lenB := len(a), len(b)

	if lenA != lenB {
		return -1, errors.New("distance mismatch")
	}
	dist := 0
	for i := 0; i < lenA; i++ {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist, nil
}
