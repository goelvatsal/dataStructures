package main

import (
	"testing"
)

func TestParametersFunction(t *testing.T) {
	num1, num2 := 10, 15
	if swap(&num1, &num2); num1 != 15 || num2 != 10 {
		t.Fatalf("expected num1 = %d, num2 = %d. got num1 = %d, num2 = %d.", 15, 10, num1, num2)
	}
}
