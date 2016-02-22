package main

import (
	"github.com/tiagofalcao/GoNotebook/manager/test"
	"testing"
)

func TestSample(t *testing.T) {
	input := "sample.in"
	output := "sample.out"

	test.Test(t, input, runCase, output, nil)
}

func BenchmarkSample(b *testing.B) {
	input := "sample.in"

	test.Benchmark(b, input, runCase)
}
