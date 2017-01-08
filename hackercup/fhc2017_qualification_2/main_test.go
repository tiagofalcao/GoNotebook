package main

import (
	"testing"

	"github.com/tiagofalcao/GoNotebook/hackercup/manager/test"
)

func TestSample(t *testing.T) {
	input := "sample.in"
	output := "sample.out"

	test.Test(t, input, runCase, output)
}

func TestFHC(t *testing.T) {
	input := "test.in"
	output := "test.out"

	test.Test(t, input, runCase, output)
}
