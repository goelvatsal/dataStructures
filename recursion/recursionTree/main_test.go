package main

import "testing"

func TestRecursionTree(t *testing.T) {
	expectedOpt := "3 2 1 1 2 1 1"
	if actualOpt := recursionTree(3); actualOpt != expectedOpt {
		t.Fatalf("Expected %s, got %s", expectedOpt, actualOpt)
	}
}
