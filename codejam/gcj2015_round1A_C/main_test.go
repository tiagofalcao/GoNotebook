package main

import (
	"github.com/tiagofalcao/GoNotebook/codejam/test"
	"testing"
)

func TestSample(t *testing.T) {
	input := "C-sample.in"
	output := "C-sample.out"

	test.Test(t, input, runCase, output)
}

func TestSmall0(t *testing.T) {
	input := "C-large-practice.in"
	output := "C-large-practice.out"

	test.Test(t, input, runCase, output)
}

func TestLarge(t *testing.T) {
	input := "C-large-practice.in"
	output := "C-large-practice.out"

	test.Test(t, input, runCase, output)
}
