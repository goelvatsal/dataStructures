package insertSort

import (
	append2 "dataStructures/arrays/arraysADT/append"
	"dataStructures/arrays/arraysADT/binSearch"
	delete2 "dataStructures/arrays/arraysADT/delete"
	"dataStructures/arrays/arraysADT/display"
	"dataStructures/arrays/arraysADT/get"
	"dataStructures/arrays/arraysADT/insert"
	"dataStructures/arrays/arraysADT/linSearch"
	maxMin "dataStructures/arrays/arraysADT/max-min"
	"dataStructures/arrays/arraysADT/reverse"
	"dataStructures/arrays/arraysADT/set"
	"dataStructures/arrays/arraysADT/sum"
)

type Array struct {
	reverse.Array
}

func (arr *Array) insertSort(x int) string {
	if arr.Length == arr.Size {
		return "Cannot insert element!"
	}

	i := arr.Length - 1
	for i >= 0 && arr.A[i] > x {
		arr.A[i+1] = arr.A[i]
		i--
	}
	arr.A[i+1] = x
	arr.Length++
	return arr.Display()
}

type ImplS1A struct{}

func (s1 ImplS1A) arraysInsertSorter(base [20]int, x int) string {
	var i int
	for _, v := range base {
		if v != 0 {
			i++
		} else {
			break
		}
	}

	arr1 := Array{reverse.Array{sum.Array{maxMin.Array{set.Array{get.Array{binSearch.Array{linSearch.Array{delete2.Array{insert.Array{append2.Array{display.Array{base, 20, i}}}}}}}}}}}}
	return arr1.insertSort(x)
}
