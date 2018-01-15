// Package hamming implements the hamming distance between DNA strings.
package hamming

import (
	"errors"
	"fmt"
)

// Distance returns the hamming distance between.
// The distance only is defined with equal length strings, the function returns error otherwise.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New(
			fmt.Sprintf("Sequences with different length, %d != %d.", len(a), len(b)))
	}

	d := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			d += 1
		}
	}
	return d, nil
}
