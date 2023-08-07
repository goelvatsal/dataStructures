package main

import "fmt"

func recursiveSumOfN(n int) int {
	if n == 0 {
		return 0
	}
	return recursiveSumOfN(n-1) + n
}

func ISumOfN(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}

func main() {
	//r := recursiveSumOfN(5)
	r := ISumOfN(5)
	fmt.Println(r)
}
