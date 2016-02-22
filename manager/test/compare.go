package test

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"
	"github.com/tiagofalcao/GoNotebook/diff"
)

func CompareTestFail(output, expected []string, t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	fmt.Fprintf(w, "Differences found\n%-35s | %-35s\n", "*** Output ***", "*** Expected ***")
	diff.PrintLinesCompare(w, output, expected)
	w.Flush()
	t.Errorf(b.String())
}


func CompareTestDefault(output, expected []string, t *testing.T) bool {
	if len(output) != len(expected) {
		CompareTestFail(output, expected, t)
		return false
	}

  for i:=0; i<len(output); i++ {
		if output[i] != expected[i] {
			CompareTestFail(output, expected, t)
			return false
		}
  }

	return true
}
