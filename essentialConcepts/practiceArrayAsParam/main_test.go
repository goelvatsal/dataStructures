package main

import (
	"reflect"
	"testing"
)

func TestArrayAsParam(t *testing.T) {
	size := 7
	ptr := fun(size)
	expectedOut := []int{1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(ptr, expectedOut) {
		t.Fatalf("expected %v, got %v", expectedOut, ptr)
	}

}
