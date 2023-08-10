package main

import "testing"

func TestRecursiveSumOfN(t *testing.T) {
	var actualOut int
	expectedOut := 15

	if actualOut = recursiveSumOfN(5); actualOut != expectedOut {
		t.Fatalf("Expected %d, got %d", expectedOut, actualOut)
	}

	if actualOut = ISumOfN(5); actualOut != expectedOut {
		t.Fatalf("Expected %d, got %d", expectedOut, actualOut)
	}
}
