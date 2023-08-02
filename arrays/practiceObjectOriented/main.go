package main

type Rectangle struct {
	length int
	width  int
}

func Monolithic(length, width int) (int, int) {
	r := Rectangle{length, width}

	area := r.findArea()
	peri := r.findPerimeter()
	return area, peri
}

func (r Rectangle) findArea() int {
	return r.length * r.width
}

func (r Rectangle) findPerimeter() int {
	return 2 * (r.length + r.width)
}
