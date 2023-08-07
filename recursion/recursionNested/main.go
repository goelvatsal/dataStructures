package main

import "fmt"

func nestedRecursion(n int) int {
	if n > 100 {
		return n - 10
	} else {
		return nestedRecursion(nestedRecursion(n + 11))
	}
}

func main() {
	r := nestedRecursion(95)
	fmt.Printf("%d\n", r)
}
