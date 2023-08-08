package main

import (
	"testing"
)

func TestRecursionTaylor(t *testing.T) {
	var expectedOut = 2.7182818284590455

	if actualOut := recursionTaylorSeries(1, 17); actualOut != expectedOut {
		t.Logf("Expected %g, got %g", expectedOut, actualOut)
	}
}
