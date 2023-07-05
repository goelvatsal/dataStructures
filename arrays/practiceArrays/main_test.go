package main

import (
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {
	result := arrayPractice(2)
	if !reflect.DeepEqual(result, []int{2, 0}) {
		t.Fatalf("Result was incorrect, got: %d, want: %v.\n", result, []int{2, 0})
	}
}
