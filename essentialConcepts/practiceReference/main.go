package main

import "fmt"

func referencePractice() int {
	a := 10
	r := &a

	var b = 30
	r = &b

	fmt.Println("A: ", a)
	fmt.Println("R: ", *r)

	return *r
}
