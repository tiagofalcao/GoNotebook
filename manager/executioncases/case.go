package executioncases

import (
	"bufio"
	"io"
)

type CasePrint func(output io.Writer, caseNum uint64, value string)

type ExecutionCase func(input *bufio.Reader, inputLock chan bool) string

type caseOutput struct {
	value   string // The result
	caseNum uint64 // The case number
	index   int    // The index of the item in the heap.
}

func CasePrintDefault(output io.Writer, caseNum uint64, value string) {
	io.WriteString(output, value)
}
