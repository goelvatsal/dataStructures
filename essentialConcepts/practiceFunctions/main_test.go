package main

import (
	"testing"
)

func TestFunctions(t *testing.T) {
	num1, num2 := 10, 15
	var sum int

	sum = add(num1, num2)

	if sum != 25 {
		t.Fatalf("Got %d, expected %d.", sum, 25)
	}
}
