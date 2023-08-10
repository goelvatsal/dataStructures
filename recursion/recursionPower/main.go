package main

import "fmt"

func power(m, n int) int {
	if n == 0 {
		return 1
	} else {
		return power(m, n-1) * m
	}
}

func power1(m, n int) int {
	if n == 0 {
		return 1
	} else if n%2 == 0 {
		return power1(m*m, n/2)
	}
	return m * power1(m*m, (n-1)/2)
}

func main() {
	r := power(2, 9)
	fmt.Println(r)

	r = power1(2, 9)
	fmt.Println(r)
}
