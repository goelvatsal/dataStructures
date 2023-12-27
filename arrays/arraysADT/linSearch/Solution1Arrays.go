package linSearch

import (
	append2 "dataStructures/arrays/arraysADT/append"
	delete2 "dataStructures/arrays/arraysADT/delete"
	"dataStructures/arrays/arraysADT/display"
	"dataStructures/arrays/arraysADT/insert"
)

type Array struct {
	delete2.Array
}

func (arr Array) LinSearch(key int) int {
	for i := 0; i < arr.Length; i++ {
		if key == arr.A[i] {
			return i
		}
	}
	return -1
}

type Impl1SA struct{}

func (s1 Impl1SA) ArrayLinSearcher(init [20]int, key int) int {
	var i int
	for _, v := range init {
		if v != 0 {
			i++
		}
	}

	arr := Array{delete2.Array{insert.Array{append2.Array{display.Array{init, 20, i}}}}}
	return arr.LinSearch(key)
}
