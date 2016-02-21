package test

import (
	"bufio"
	"github.com/tiagofalcao/GoNotebook/diff"
	"os"
	"testing"
)

import manager "github.com/tiagofalcao/GoNotebook/manager/executioncases"

func callbackErr(base, received []byte, value interface{}) {
	t := value.(*testing.T)
	t.Errorf("Differ %s and %s", string(base), string(received))
}

func Test(t *testing.T, input string, caseTask manager.ExecutionCase, print manager.CasePrint, output string) {
	i, err := os.Open(input)
	if err != nil {
		t.Errorf("Can't open %s", input)
	}
	b, err := os.Open(output)
	if err != nil {
		t.Errorf("Can't open %s", output)
	}
	o := diff.NewBuf(b, callbackErr, t)

	manager.NewExecutionManagerIO(caseTask, print, bufio.NewReader(i), o).WaitEnd()

	i.Close()
	b.Close()
}
