package main

import "fmt"

type Rectangle struct {
	length  int
	breadth int
}

func structsPractice() Rectangle {
	r1 := Rectangle{10, 5}

	r1.length = 15
	r1.breadth = 7

	fmt.Println("Length: ", r1.length)
	fmt.Println("Breadth (Width): ", r1.breadth)

	return r1
}
