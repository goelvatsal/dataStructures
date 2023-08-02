package main

import "testing"

func TestMonolithic(t *testing.T) {
	area, peri := Monolithic(10, 5)

	if area != 50 || peri != 30 {
		t.Fatalf("Expected area and perimeter to be %d, %d. Got area and perimeter as %d, %d.", 50, 30, area, peri)
	}
}
