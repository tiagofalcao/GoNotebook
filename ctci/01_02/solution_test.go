package solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		description string
		input       string
		expect      string
	}{
		{"Simple", "abcd", "dcba"},
		{"Empty", "", ""},
		{"Symbols", `_\/,"'.;`, `;.'",/\_`},
		{"Palindrome", "abcdeedcba", "abcdeedcba"},
	}

	for _, c := range cases {
		r := Reverse(c.input)
		if r != c.expect {
			t.Errorf("Reverse(%v) [%v]: expected %v, got %v", c.input, c.description, c.expect, r)
		}
	}
}
