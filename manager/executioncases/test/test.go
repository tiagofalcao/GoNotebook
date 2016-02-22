package test

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/tiagofalcao/GoNotebook/diff"
	"os"
	"testing"
)

import manager "github.com/tiagofalcao/GoNotebook/manager/executioncases"

type CompareTest func(output, expected []string, t *testing.T) bool

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

func Test(t *testing.T, input string, caseTask manager.ExecutionCase, print manager.CasePrint, output string, cmp CompareTest) {
	i, err := os.Open(input)
	if err != nil {
		t.Errorf("Can't open %s", input)
	}
	b, err := os.Open(output)
	if err != nil {
		t.Errorf("Can't open %s", output)
	}
	if cmp == nil {
		cmp = CompareTestDefault
	}
	test := func(output, expected []string, value interface{}) bool {
			t := value.(*testing.T)
			return cmp(output, expected, t)
	}
	o := diff.NewBuf(b, test, t)

	manager.NewExecutionManagerIO(caseTask, print, bufio.NewReader(i), o).WaitEnd()

	i.Close()
	b.Close()
}
