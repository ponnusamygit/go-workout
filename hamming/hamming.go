package hamming

import (
	"errors"
)

// Distance this will return hamming difference
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Given strings are not matching")
	}
	difference := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			difference++
		}
	}
	return difference, nil
}
