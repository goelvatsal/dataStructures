package main

import (
	"reflect"
	"testing"
)

func TestStructs(t *testing.T) {
	result := structsPractice()
	if !reflect.DeepEqual(result, Rectangle{15, 7}) {
		t.Fatalf("Result was incorrect, got: %d, want: %v.\n", result, Rectangle{15, 7})
	}
}
