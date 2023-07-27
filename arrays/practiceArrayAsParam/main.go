package main

func fun(size int) []int {
	var p []int
	for i := 0; i < size; i++ {
		p = append(p, i+1)
	}
	return p
}
