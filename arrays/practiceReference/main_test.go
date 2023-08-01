package main

import "testing"

func TestReference(t *testing.T) {
	result := referencePractice()

	if result != 30 {
		t.Fatalf("Expected 30, got %d", result)
	}
}
