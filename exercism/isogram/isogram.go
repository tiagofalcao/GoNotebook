// Package isogram implements the functions to verify isograms.
package isogram

// Similar problem is presented in CtCI book.
import solution "github.com/tiagofalcao/GoNotebook/ctci/01_01"

// IsIsogram returns true if the string s is a isogram.
func IsIsogram(s string) bool {
	return solution.UniqueLetters(s)
}
