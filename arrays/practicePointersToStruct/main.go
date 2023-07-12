package main

import "fmt"

type Rectangle struct {
	length  int
	breadth int
}

func pointersPractice() []int {
	var p *Rectangle

	p = new(Rectangle)
	p = &Rectangle{
		length:  15,
		breadth: 7,
	}

	fmt.Println("length =", p.length)
	fmt.Println("breadth =", p.breadth)

	return []int{p.length, p.breadth}
}
