package main

import (
	"github.com/tiagofalcao/GoNotebook/codejam/manager/test"
	"testing"
)

func TestSample(t *testing.T) {
	input := "sample.in"
	output := "sample.out"

	test.Test(t, input, runCase, output)
}

func TestSmall0(t *testing.T) {
	input := "small-practice.in"
	output := "small-practice.out"

	test.Test(t, input, runCase, output)
}

func TestLarge(t *testing.T) {
	input := "large-practice.in"
	output := "large-practice.out"

	test.Test(t, input, runCase, output)
}

func BenchmarkLarge(b *testing.B) {
	input := "large-practice.in"

	test.Benchmark(b, input, runCase)
}
