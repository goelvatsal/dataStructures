package main

func arrayPractice(n int) []int {
	a := make([]int, n)
	a[0] = 2

	var out []int

	for _, j := range a {
		out = append(out, j)
	}
	return out
}
