package main

import "fmt"

var p, f float64 = 1, 1

func recursionTaylorSeries(x, n float64) float64 {
	if n == 0 {
		return 1
	} else {
		r := recursionTaylorSeries(x, n-1)
		p *= x
		f *= n
		return r + p/f
	}
}

func main() {
	//fmt.Println(recursionTaylorSeries(1, 17))
	fmt.Println(recursionTaylorSeries(10, 45))
}
