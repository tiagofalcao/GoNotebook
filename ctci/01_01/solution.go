package solution

import (
	"bufio"
	"fmt"
	"unicode"

	"github.com/tiagofalcao/GoNotebook/log"
)

func UniqueRunes(str string) bool {
	m := make(map[rune]bool)
	for i, c := range str {
		_, ok := m[c]
		if ok {
			log.Debug.Printf("Duplicated %v at %d column", c, i)
			return false
		}
		m[c] = true
	}
	return true
}

// UniqueLetter returns true if the string is composed by unique letters, ignoring the case.
// Symbols can be repeated and are ignored.
func UniqueLetters(str string) bool {
	m := make(map[rune]bool)
	for i, c := range str {
		if !unicode.IsLetter(c) {
			continue
		}

		c = unicode.ToLower(c)
		_, ok := m[c]
		if ok {
			log.Debug.Printf("Duplicated %v at %d column", c, i)
			return false
		}
		m[c] = true
	}
	return true
}

func RunCase(input *bufio.Reader, inputLock chan bool) (result string) {
	str, err := input.ReadString('\n')
	inputLock <- true

	if err != nil {
		return " error"
	}

	return fmt.Sprintf(" %v\n", UniqueRunes(str))
}
