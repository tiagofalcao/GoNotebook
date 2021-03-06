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
	output := "B-small-attempt0.out"

	test.Test(t, input, runCase, output)
}

func TestSmall1(t *testing.T) {
	input := "small-attempt1.in"
	output := "B-small-attempt1.out"

	test.Test(t, input, runCase, output)
}

func TestSmall2(t *testing.T) {
	input := "small-attempt2.in"
	output := "B-small-attempt2.out"

	test.Test(t, input, runCase, output)
}

func TestSmall3(t *testing.T) {
	input := "small-attempt3.in"
	output := "B-small-attempt3.out"

	test.Test(t, input, runCase, output)
}

func TestLargePractice(t *testing.T) {
	input := "small-practice.in"
	output := "B-small-practice.out"

	test.Test(t, input, runCase, output)
}

func TestSmallPractice(t *testing.T) {
	input := "large-practice.in"
	output := "B-large-practice.out"

	test.Test(t, input, runCase, output)
}


func BenchmarkLarge(b *testing.B) {
	input := "large-practice.in"

	test.Benchmark(b, input, runCase)
}
