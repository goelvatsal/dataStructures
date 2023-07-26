package main

func swap(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
