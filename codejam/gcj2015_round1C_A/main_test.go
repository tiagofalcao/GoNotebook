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
	input := "A-small-attempt0.in"
	output := "A-small-attempt0.out"

	test.Test(t, input, runCase, output)
}

func TestSmall1(t *testing.T) {
	input := "A-small-attempt1.in"
	output := "A-small-attempt1.out"

	test.Test(t, input, runCase, output)
}

func TestLarge(t *testing.T) {
	input := "A-large.in"
	output := "A-large.out"

	test.Test(t, input, runCase, output)
}
