package main

import "fmt"

var slc []int

func fibInit() {
	slc = make([]int, 20)
	for i := range slc {
		slc[i] = -1
	}
}

func recursiveFib(n int) int {
	if n <= 1 {
		slc[n] = n
		return n
	} else {
		if slc[n-2] == -1 {
			slc[n-2] = recursiveFib(n - 2)
		}
		if slc[n-1] == -1 {
			slc[n-1] = recursiveFib(n - 1)
		}
		slc[n] = slc[n-2] + slc[n-1]
		return slc[n-2] + slc[n-1]
	}
}

func main() {
	fibInit()
	fmt.Println(recursiveFib(10))
}
