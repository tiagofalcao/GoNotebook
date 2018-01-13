// Package bob produces the lackadaisical teenager responses in conversartion.
package bob

import "unicode"

// Hey produces the response for one remark input.
func Hey(remark string) string {
	lower := uint(0)
	upper := uint(0)
	last := ' '
	for _, r := range remark {
		if unicode.IsSpace(r) {
			continue
		}
		last = r
		if unicode.IsLower(r) {
			lower += 1
		} else if unicode.IsUpper(r) {
			upper += 1
		}
	}

	if last == '?' {
		if upper > 0 && lower == 0 {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}

	if upper > 0 && lower == 0 {
		return "Whoa, chill out!"
	}

	if unicode.IsSpace(last) {
		return "Fine. Be that way!"
	}

	return "Whatever."
}
