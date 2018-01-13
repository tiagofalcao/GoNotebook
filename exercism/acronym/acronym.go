// Package acronym offers functions to abbreviate one string.
package acronym

import "unicode"

// Abbreviate returns a string with initial letters of all words in upper-case.
func Abbreviate(s string) string {
	getchar := true
	letters := make([]rune, 0, 10)
	for _, r := range s {
		if !unicode.IsLetter(r) {
			getchar = true
			continue
		}
		if !getchar {
			continue
		}
		getchar = false
		letters = append(letters, unicode.ToUpper(r))
	}
	return string(letters)
}
