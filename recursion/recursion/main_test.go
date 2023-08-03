package main

import "testing"

func TestRecursion(t *testing.T) {
	expectedOpt := "3 2 1"

	if actualOut := recursion(3); actualOut != expectedOpt {
		t.Fatalf("Expected %q, got %q", expectedOpt, actualOut)
	}
}
