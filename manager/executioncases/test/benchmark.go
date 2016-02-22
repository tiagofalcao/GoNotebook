package test

import (
	"bufio"
	"os"
	"testing"
)

import manager "github.com/tiagofalcao/GoNotebook/manager/executioncases"
import test "github.com/tiagofalcao/GoNotebook/manager/test"

func Benchmark(b *testing.B, input string, caseTask manager.ExecutionCase, print manager.CasePrint) {
  	i, err := os.Open(input)
  	if err != nil {
  		b.Errorf("Can't open %s", input)
  	}
    var o test.NilWriter
    b.ResetTimer()

    for x := 0; x < b.N; x++ {
			manager.NewExecutionManagerIO(caseTask, print, bufio.NewReader(i), o).WaitEnd()
    }

    i.Close()
}
