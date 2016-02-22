package diff

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Callback is called when bytes differ
type Compare func(output, expected []string, v interface{}) bool

// Buf simulate a io.Writer to check differences
type Buf struct {
	base     *os.File
	scanner  *bufio.Scanner
	compare  Compare
	value    interface{}
}

// NewBuf create a new Buf
func NewBuf(b *os.File, c Compare, v interface{}) Buf {
	return Buf{b, bufio.NewScanner(b), c, v}
}

// Write simulate write and check bytes
func (d Buf) Write(p []byte) (n int, err error) {
	lines := strings.Split(string(p), "\n")

	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}

	answer := make([]string, 0, len(lines))
	i := 0

	for d.scanner.Scan() {
		answer = append(answer, d.scanner.Text())
		i += 1
		if i == len(lines) {
			break
		}
	}

	if d.compare(lines, answer, d.value) {
		return len(p), nil
	}

	return len(p), fmt.Errorf("Output differ")

}

//Close check end of buffer
func (d Buf) Close() error {
	return nil
}
