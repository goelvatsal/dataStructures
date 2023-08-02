package main

import (
	"testing"
)

func TestStructureAsParam(t *testing.T) {
	ptr := fun()
	if ptr.length != 15 {
		t.Fatalf("expected length to be %d, got %d", 15, ptr.length)
	}

	if ptr.width != 7 {
		t.Fatalf("expected width to be %d, got %d", 7, ptr.width)
	}
}
