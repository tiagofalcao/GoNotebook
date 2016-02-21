package main

import (
	"github.com/tiagofalcao/GoNotebook/codejam/manager/test"
	"testing"
)

func TestSample(t *testing.T) {
	input := "A-sample.in"
	output := "A-sample.out"

	test.Test(t, input, runCase, output)
}

func TestSmall0(t *testing.T) {
	input := "A-small-practice.in"
	output := "A-small-practice.out"

	test.Test(t, input, runCase, output)
}

func TestLarge(t *testing.T) {
	input := "A-large-practice.in"
	output := "A-large-practice.out"

	test.Test(t, input, runCase, output)
}
