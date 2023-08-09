package main

import "fmt"

var s float64

func recursionHornersRule(x, n float64) float64 {
	if n == 0 {
		return s
	}

	s = 1 + x*s/n
	return recursionHornersRule(x, n-1)
}

func main() {
	fmt.Println(recursionHornersRule(2, 25))
}
