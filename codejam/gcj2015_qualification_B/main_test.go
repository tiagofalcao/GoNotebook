package main

import (
	"github.com/tiagofalcao/GoNotebook/codejam/test"
	"testing"
)

func TestSample(t *testing.T) {
	input := "B-sample.in"
	output := "B-sample.out"

	test.Test(t, input, runCase, output)
}

func TestSmall0(t *testing.T) {
	input := "B-small-attempt0.in"
	output := "B-small-attempt0.out"

	test.Test(t, input, runCase, output)
}

func TestSmall1(t *testing.T) {
	input := "B-small-attempt1.in"
	output := "B-small-attempt1.out"

	test.Test(t, input, runCase, output)
}

func TestSmall2(t *testing.T) {
	input := "B-small-attempt2.in"
	output := "B-small-attempt2.out"

	test.Test(t, input, runCase, output)
}

func TestSmall3(t *testing.T) {
	input := "B-small-attempt3.in"
	output := "B-small-attempt3.out"

	test.Test(t, input, runCase, output)
}

func TestLargePractice(t *testing.T) {
	input := "B-small-practice.in"
	output := "B-small-practice.out"

	test.Test(t, input, runCase, output)
}

func TestSmallPractice(t *testing.T) {
	input := "B-large-practice.in"
	output := "B-large-practice.out"

	test.Test(t, input, runCase, output)
}
