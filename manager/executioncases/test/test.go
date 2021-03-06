package test

import (
	"bufio"
	"github.com/tiagofalcao/GoNotebook/diff"
	"os"
	"testing"
)

import manager "github.com/tiagofalcao/GoNotebook/manager/executioncases"
import test "github.com/tiagofalcao/GoNotebook/manager/test"

func Test(t *testing.T, input string, caseTask manager.ExecutionCase, print manager.CasePrint, output string, cmp test.CompareTest) {
	i, err := os.Open(input)
	if err != nil {
		t.Errorf("Can't open %s", input)
	}
	b, err := os.Open(output)
	if err != nil {
		t.Errorf("Can't open %s", output)
	}
	if cmp == nil {
		cmp = test.CompareTestDefault
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
