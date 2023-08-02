package main

func Monolithic(length, width int) (int, int) {
	var area, peri int

	area = length * width
	peri = 2 * (length + width)
	return area, peri
}
