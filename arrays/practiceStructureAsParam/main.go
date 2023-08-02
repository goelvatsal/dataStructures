package main

type Rectangle struct {
	length int
	width  int
}

func fun() *Rectangle {
	p := new(Rectangle)

	p.length = 15
	p.width = 7

	return p
}
