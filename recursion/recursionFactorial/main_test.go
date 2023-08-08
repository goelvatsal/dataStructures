package main

import "testing"

func TestRecursiveFactorial(t *testing.T) {
	expectedOut := 120

	if actualOut := recursionFactorial(5); actualOut != expectedOut {
		t.Fatalf("Expected %d, got %d", expectedOut, actualOut)
	}

	if actualOut := iterativeFactorial(5); actualOut != expectedOut {
		t.Fatalf("Expected %d, got %d", expectedOut, actualOut)
	}
}
