package binSearch

import (
	append2 "dataStructures/arrays/arraysADT/append"
	delete2 "dataStructures/arrays/arraysADT/delete"
	"dataStructures/arrays/arraysADT/display"
	"dataStructures/arrays/arraysADT/insert"
	"dataStructures/arrays/arraysADT/linSearch"
)

type Array struct {
	linSearch.Array
}

func (arr Array) BinSearch(key int) int {
	var low, mid, high int
	high = arr.Length - 1

	for low <= high {
		mid = (low + high) / 2
		if key == arr.A[mid] {
			return mid
		} else if key < arr.A[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

type ImplS1A struct{}

func (s1 ImplS1A) arrayBinSearcher(init [20]int, key int) int {
	var i int
	for _, v := range init {
		if v != 0 {
			i++
		}
	}

	arr := Array{linSearch.Array{delete2.Array{insert.Array{append2.Array{display.Array{init, 20, i}}}}}}
	return arr.BinSearch(key)
}
