package negSort

import (
	append2 "dataStructures/arrays/arraysADT/append"
	"dataStructures/arrays/arraysADT/binSearch"
	delete2 "dataStructures/arrays/arraysADT/delete"
	"dataStructures/arrays/arraysADT/display"
	"dataStructures/arrays/arraysADT/get"
	"dataStructures/arrays/arraysADT/insert"
	"dataStructures/arrays/arraysADT/insertSort"
	"dataStructures/arrays/arraysADT/isSorted"
	"dataStructures/arrays/arraysADT/linSearch"
	maxMin "dataStructures/arrays/arraysADT/max-min"
	"dataStructures/arrays/arraysADT/reverse"
	"dataStructures/arrays/arraysADT/set"
	"dataStructures/arrays/arraysADT/sum"
)

type Array struct {
	isSorted.Array
}

func (arr *Array) negSort() string {
	i, j := 0, arr.Length-1

	for i < j {
		for arr.A[i] < 0 && i < j {
			i++
		}
		for arr.A[i] >= 0 && i < j {
			j--
		}
		if i < j {
			arr.A[i], arr.A[j] = arr.A[j], arr.A[i]
			i++
			j--
		}
	}
	return arr.Display()
}

type ImplS1A struct{}

func (s1 ImplS1A) arraysNegSorter(base [20]int) string {
	var i int
	for _, v := range base {
		if v != 0 {
			i++
		} else {
			break
		}
	}

	arr := Array{isSorted.Array{insertSort.Array{reverse.Array{sum.Array{maxMin.Array{set.Array{get.Array{binSearch.Array{linSearch.Array{delete2.Array{insert.Array{append2.Array{display.Array{base, 20, i}}}}}}}}}}}}}}
	return arr.negSort()
}
