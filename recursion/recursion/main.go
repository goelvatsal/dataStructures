package main

import "fmt"

func fun1(n int) {
	if n > 0 {
		fmt.Printf("%d ", n)
		fun1(n - 1)
	}
}

func main() {
	fun1(3)
	fmt.Println()
}
