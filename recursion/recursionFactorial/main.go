package main

import "fmt"

func recursionFactorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return recursionFactorial(n-1) * n
	}
}

func iterativeFactorial(n int) int {
	product := 1
	for i := 1; i <= n; i++ {
		product *= i
	}
	return product
}

func main() {
	out := recursionFactorial(5)
	fmt.Println(out)

	out = iterativeFactorial(5)
	fmt.Println(out)
}
