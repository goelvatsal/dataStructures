package main

import (
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {
	result := structsPractice()
	if !reflect.DeepEqual(result, Card{1, 0, 0}) {
		t.Fatalf("Result was incorrect, got: %d, want: %v.\n", result, Card{1, 0, 0})
	}
}
