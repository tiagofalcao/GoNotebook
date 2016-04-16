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

func TestSample2(t *testing.T) {
	input := "sample2.in"
	output := "sample2.out"

	test.Test(t, input, runCase, output)
}

func TestSmall0(t *testing.T) {
	input := "small-attempt0.in"
	output := "small-attempt0.out"

	test.Test(t, input, runCase, output)
}

func TestLargePractice(t *testing.T) {
	input := "small-practice.in"
	output := "small-practice.out"

	test.Test(t, input, runCase, output)
}

func TestSmallPractice(t *testing.T) {
	input := "large-practice.in"
	output := "large-practice.out"

	test.Test(t, input, runCase, output)
}

func BenchmarkLarge(b *testing.B) {
	input := "large-practice.in"

	test.Benchmark(b, input, runCase)
}
