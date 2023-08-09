package main

import "testing"

func TestRecursionFib(t *testing.T) {
	expectedOut := 55

	fibInit()
	if actualOut := recursiveFib(10); actualOut != expectedOut {
		t.Fatalf("Expected %d, got %d", expectedOut, actualOut)
	}
}
