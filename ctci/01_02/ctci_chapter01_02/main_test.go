package main

import (
	"testing"

	solution "github.com/tiagofalcao/GoNotebook/ctci/01_02"
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

func BenchmarkSolutionA(b *testing.B) {
	input := "benchmark.in"

	solution.Reverse = solution.ReverseA
	test.Benchmark(b, input, solution.RunCase)
}

func BenchmarkSolutionB(b *testing.B) {
	input := "benchmark.in"

	solution.Reverse = solution.ReverseB
	test.Benchmark(b, input, solution.RunCase)
}
