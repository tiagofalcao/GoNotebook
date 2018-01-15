// Package twofer implements the two-fer string format.
package twofer

import "fmt"

// ShareWith returns the two-fer string for given name n.
func ShareWith(n string) string {
	if n == "" {
		n = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", n)
}
