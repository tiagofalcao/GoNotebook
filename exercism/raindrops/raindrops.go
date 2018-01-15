// Package raindrops implements the respective number conversion to string.
package raindrops

import "strconv"

// Convert converts the number i to string using raindrops.
func Convert(i int) string {
	var r string
	if i%3 == 0 {
		r += "Pling"
	}
	if i%5 == 0 {
		r += "Plang"
	}
	if i%7 == 0 {
		r += "Plong"
	}
	if len(r) == 0 {
		return strconv.Itoa(i)
	}
	return r
}
