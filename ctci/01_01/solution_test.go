package solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		description string
		input       string
		expect      bool
	}{
		{"Simple", "abcd", true},
		{"Empty", "", true},
		{"Symbols", `_\/,"'.;`, true},
		{"Repeated", "abca", false},
		{"Uniform", "aaaa", false},
	}

	for _, c := range cases {
		r := UniqueRunes(c.input)
		if r != c.expect {
			t.Errorf("UniqueRunes(%v): expected %v, got %v", c.expect, r)
		}
	}
}
