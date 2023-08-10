package main

import (
	"strings"
	"testing"
)

func TestRecursionIndirect(t *testing.T) {
	expectedOut := "20 19 9 8 4 3 1"
	var actualOut string

	if a(20, &actualOut); strings.TrimSpace(actualOut) != expectedOut {
		t.Fatalf("Expected %s, got %v", expectedOut, actualOut)
	}
}
