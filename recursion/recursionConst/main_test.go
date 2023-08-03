package main

import "testing"

func TestRecursionConst(t *testing.T) {
	expectedOut := 25
	if actualOut := recursionConst(5); actualOut != expectedOut {
		t.Fatalf("Expected %d, got %d", expectedOut, actualOut)
	}
}
