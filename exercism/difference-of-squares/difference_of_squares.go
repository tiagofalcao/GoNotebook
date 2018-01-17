// Package diffsquares implements the sum of squares and square of sums.
package diffsquares

// SquareOfSums returns the square of sums from 1 to n.
func SquareOfSums(n int) int {
	s := int(n)
	s *= s + 1
	s /= 2
	return s * s
}

// SumOfSquares returns the sums of squares from 1 to n.
func SumOfSquares(n int) int {
	s := int(n)
	s = 2*s*s*s + 3*s*s + s
	s /= 6
	return s

}

// Difference returns the difference between square of sums and sum of squares from 1 to n.
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
