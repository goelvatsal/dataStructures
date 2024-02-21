package sum

import (
	append2 "dataStructures/arrays/arraysADT/append"
	"dataStructures/arrays/arraysADT/binSearch"
	delete2 "dataStructures/arrays/arraysADT/delete"
	"dataStructures/arrays/arraysADT/display"
	"dataStructures/arrays/arraysADT/get"
	"dataStructures/arrays/arraysADT/insert"
	"dataStructures/arrays/arraysADT/linSearch"
	maxMin "dataStructures/arrays/arraysADT/max-min"
	"dataStructures/arrays/arraysADT/set"
)

type Array struct {
	maxMin.Array
}

func (arr Array) Sum() int {
	total := 0
	for i := 0; i < arr.Length; i++ {
		total += arr.A[i]
	}
	return total
}

type ImplS1A struct{}

func (s1 ImplS1A) arraySummer(base [20]int) int {
	var i int
	for _, v := range base {
		if v != 0 {
			i++
		} else {
			break
		}
	}

	arr := Array{maxMin.Array{set.Array{get.Array{binSearch.Array{linSearch.Array{delete2.Array{insert.Array{append2.Array{display.Array{base, 20, i}}}}}}}}}}
	return arr.Sum()
}
