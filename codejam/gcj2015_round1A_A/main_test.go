package main

import (
	"github.com/tiagofalcao/GoNotebook/codejam/manager"
	"github.com/tiagofalcao/GoNotebook/diff"
	"os"
	"testing"
)

func callbackErr(base, received []byte, value interface{}) {
	t := value.(*testing.T)
	t.Errorf("Differ %s and %s", string(base), string(received))
}

func defaultTest(t *testing.T, input, output string) {
	i, err := os.Open(input)
	if err != nil {
		t.Errorf("Can't open %s", input)
	}
	b, err := os.Open(output)
	if err != nil {
		t.Errorf("Can't open %s", output)
	}
	o := diff.NewBuf(b, callbackErr, t)

	manager.NewGCJManagerIO(runCase, i, o).WaitEnd()

	o.Close()
}

func TestLarge(t *testing.T) {
	input := "A-large.in"
	output := "A-large.out"

	defaultTest(t, input, output)
}
