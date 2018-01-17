// Package luhn implements the Luhn algorithm.
package luhn

import "unicode"

// Valid verifies if one string is Luhn valid.
func Valid(c string) bool {
	// Get the digits, safe for utf-8
	d := make([]uint8, 0, 20)
	for _, r := range c {
		if !unicode.IsDigit(r) {
			if !unicode.IsSpace(r) {
				return false
			}
			continue
		}
		d = append(d, uint8(r)-'0')
	}

	if len(d) <= 1 {
		return false
	}

	// Luhn's sum computation
	s := uint(0)
	m := len(d)%2 == 0
	for _, v := range d {
		if m {
			v *= 2
			if v > 9 {
				v -= 9
			}
		}
		s += uint(v)
		m = !m
	}
	return s%10 == 0
}
