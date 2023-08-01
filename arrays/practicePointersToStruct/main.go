package main

import "fmt"

type Rectangle struct {
	length  int
	breadth int
}

func pointersPractice(length, width int) []int {
	var p *Rectangle

	p = new(Rectangle)
	p = &Rectangle{
		length,
		width,
	}

	fmt.Println("length =", p.length)
	fmt.Println("breadth =", p.breadth)

	return []int{p.length, p.breadth}
}
