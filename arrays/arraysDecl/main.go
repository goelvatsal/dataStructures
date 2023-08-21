package main

import "fmt"

func main() {
	var A [5]int

	for i := range A {
		fmt.Println(&A[i])
	}
}
