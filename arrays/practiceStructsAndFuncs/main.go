package main

type Rectangle struct {
	length int
	width  int
}

func Monolithic(length, width int) (int, int) {
	var area, peri int

	var r Rectangle
	initialize(&r, length, width)

	area = findArea(r)
	peri = findPerimeter(r)
	return area, peri
}

func findArea(r Rectangle) int {
	return r.length * r.width
}

func findPerimeter(r Rectangle) int {
	return 2 * (r.length + r.width)
}

func initialize(r *Rectangle, length, width int) {
	r.length = length
	r.width = width
}
