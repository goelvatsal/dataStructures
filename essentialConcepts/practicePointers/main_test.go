package main

import (
	"fmt"
	"testing"
)

func TestPointers(t *testing.T) {
	result := pointerPractice()

	for _, j := range result {
		if j != 8 {
			fmt.Println(j)
			t.Fatalf("Expected %d, got %d", 8, j)
		}
	}
}
