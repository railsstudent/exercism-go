package hamming

import "errors"

// Distance calculates hamming distance
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("distance mismatch")
	}
	dist := 0
	for i, c := range a {
		if c != rune(b[i]) {
			dist++
		}
	}
	return dist, nil
}
