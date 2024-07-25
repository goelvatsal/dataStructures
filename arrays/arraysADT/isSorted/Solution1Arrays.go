package isSorted

import (
	append2 "dataStructures/arrays/arraysADT/append"
	"dataStructures/arrays/arraysADT/binSearch"
	delete2 "dataStructures/arrays/arraysADT/delete"
	"dataStructures/arrays/arraysADT/display"
	"dataStructures/arrays/arraysADT/get"
	"dataStructures/arrays/arraysADT/insert"
	"dataStructures/arrays/arraysADT/insertSort"
	"dataStructures/arrays/arraysADT/linSearch"
	maxMin "dataStructures/arrays/arraysADT/max-min"
	"dataStructures/arrays/arraysADT/reverse"
	"dataStructures/arrays/arraysADT/set"
	"dataStructures/arrays/arraysADT/sum"
)

type Array struct {
	insertSort.Array
}

func (arr Array) isSorted() string {
	for i := 0; i < arr.Length-1; i++ {
		if arr.A[i] > arr.A[i+1] {
			return "Not Sorted!"
		}
	}
	return "Sorted!"
}

type ImplS1A struct{}

func (s1 ImplS1A) arraysIsSorted(base [20]int) string {
	var i int
	for _, v := range base {
		if v != 0 {
			i++
		} else {
			break
		}
	}

	arr := Array{insertSort.Array{reverse.Array{sum.Array{maxMin.Array{set.Array{get.Array{binSearch.Array{linSearch.Array{delete2.Array{insert.Array{append2.Array{display.Array{base, 20, i}}}}}}}}}}}}}
	return arr.isSorted()
}
