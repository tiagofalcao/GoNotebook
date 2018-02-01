// Package grains implements functions to calculate the number of grains on a chessboard.
package grains

import "fmt"

// Square returns the value of grains on each square from 1 to 64.
func Square(i int) (uint64, error) {
	if i <= 0 || i > 64 {
		return 0, fmt.Errorf("Invalid value: %d", i)
	}
	r := uint64(1) << uint(i-1)
	return r, nil
}

// Total returns the total value of grains on a chessboard.
func Total() uint64 {
	return ^uint64(0)
}
