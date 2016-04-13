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
	input := "small-attempt0.in"
	output := "C-small-attempt0.out"

	test.Test(t, input, runCase, output)
}

func TestSmall1(t *testing.T) {
	input := "small-attempt1.in"
	output := "C-small-attempt1.out"

	test.Test(t, input, runCase, output)
}

func TestSmall2(t *testing.T) {
	input := "small-attempt2.in"
	output := "C-small-attempt2.out"

	test.Test(t, input, runCase, output)
}

func TestLarge(t *testing.T) {
	input := "large.in"
	output := "C-large.out"

	test.Test(t, input, runCase, output)
}

func TestLargePractice(t *testing.T) {
	input := "small-practice.in"
	output := "C-small-practice.out"

	test.Test(t, input, runCase, output)
}

func TestSmallPractice(t *testing.T) {
	input := "large-practice.in"
	output := "C-large-practice.out"

	test.Test(t, input, runCase, output)
}

func BenchmarkLarge(b *testing.B) {
	input := "large-practice.in"

	test.Benchmark(b, input, runCase)
}
