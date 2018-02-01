package solution

import (
	"bufio"
	"fmt"
)

var Reverse = ReverseA
var Reverses = []func(string) string{ReverseA, ReverseB}

func ReverseA(str string) string {
	s := []rune(str)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func ReverseB(str string) string {
	s := make([]rune, len(str))
	for i, r := range str {
		s[i] = r
	}
	return string(s)
}

func RunCase(input *bufio.Reader, inputLock chan bool) (result string) {
	str, err := input.ReadString('\n')
	inputLock <- true

	if err != nil {
		return " error"
	}

	return fmt.Sprintf(" %v\n", Reverse(str))
}
