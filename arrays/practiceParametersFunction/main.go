package main

func swap(x, y *int) {
	*x, *y = *y, *x
}
