package main

import "testing"

func TestRecursionNested(t *testing.T) {
	r := nestedRecursion(95)
	if r != 91 {
		t.Errorf("Expected 91, got %d", r)
	}
}
