package main

var x int

func recursionConst(n int) int {
	if n > 0 {
		x++
		return recursionConst(n-1) + x
	}
	return 0
}
