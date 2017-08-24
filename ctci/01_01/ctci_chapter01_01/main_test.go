package main

import (
	"testing"

	solution "github.com/tiagofalcao/GoNotebook/ctci/01_01"
	"github.com/tiagofalcao/GoNotebook/ctci/manager/test"
)

func TestSolution(t *testing.T) {
	input := "test.in"
	output := "test.out"

	test.Test(t, input, solution.RunCase, output)
}

func BenchmarkSolution(b *testing.B) {
	input := "benchmark.in"

	test.Benchmark(b, input, solution.RunCase)
}
