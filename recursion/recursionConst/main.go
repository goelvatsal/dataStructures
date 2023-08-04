package main

import "fmt"

var x int

func recursionConst(n int) int {
	if n > 0 {
		x++
		return recursionConst(n-1) + x
	}
	return 0
}

func main() {
	fmt.Println(recursionConst(5))
}
