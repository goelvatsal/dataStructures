package main

import "testing"

func TestMonolithic(t *testing.T) {
	area, peri := Monolithic(10, 5)
	ExpectedArea, ExpectedPeri := 50, 30

	if area != ExpectedArea || peri != ExpectedPeri {
		t.Fatalf("Expected area and perimeter to be %d, %d. Got area and perimeter as %d, %d.", ExpectedArea, ExpectedPeri, area, peri)
	}
}
