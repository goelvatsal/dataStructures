package main

import (
	"reflect"
	"testing"
)

func TestPointersStruct(t *testing.T) {
	result := pointersPractice(15, 7)
	if !reflect.DeepEqual(result, []int{15, 7}) {
		t.Fatalf("Expected {15,7}, got %v", result)
	}
}
