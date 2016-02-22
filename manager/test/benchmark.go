package test

import (
	"bufio"
	"os"
	"testing"
)

import manager "github.com/tiagofalcao/GoNotebook/manager"

func Benchmark(b *testing.B, input string, caseTask manager.Task) {
  	i, err := os.Open(input)
  	if err != nil {
  		b.Errorf("Can't open %s", input)
  	}
    var o NilWriter
    b.ResetTimer()

    for x := 0; x < b.N; x++ {
      manager.NewManagerIO(caseTask, bufio.NewReader(i), o)
    }

    i.Close()
}
