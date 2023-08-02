package main

func Monolithic(length, width int) (int, int) {
	var area, peri int

	area = findArea(length, width)
	peri = findPerimeter(length, width)
	return area, peri
}

func findArea(length, width int) int {
	return length * width
}

func findPerimeter(length, width int) int {
	return 2 * (length + width)
}
