package test

import (
	"bufio"
	"github.com/tiagofalcao/GoNotebook/diff"
	"os"
	"testing"
)

import manager "github.com/tiagofalcao/GoNotebook/manager"

type CompareTest func(output, expected []string, t *testing.T) bool

func Test(t *testing.T, input string, caseTask manager.Task, output string, cmp CompareTest) {
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

	manager.NewManagerIO(caseTask, bufio.NewReader(i), o)

	i.Close()
	b.Close()
}
