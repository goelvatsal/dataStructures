package main

import "testing"

func TestRecursionHornersRule(t *testing.T) {
	expectedOut := 7.38905609893065

	if actualOut := recursionHornersRule(2, 25); actualOut != expectedOut {
		t.Fatalf("Expected %g, got %g", expectedOut, actualOut)
	}
}
