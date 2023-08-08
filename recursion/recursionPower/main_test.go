package main

import "testing"

func TestRecursionPower(t *testing.T) {
	expectedOut := 512
	var actualOut int

	if actualOut = power(2, 9); actualOut != expectedOut {
		t.Fatalf("Expected %d, got %d", expectedOut, actualOut)
	}

	if actualOut = power1(2, 9); actualOut != expectedOut {
		t.Fatalf("Expected %d, got %d", expectedOut, actualOut)
	}
}
