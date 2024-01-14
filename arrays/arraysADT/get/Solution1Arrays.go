package get

import (
	append2 "dataStructures/arrays/arraysADT/append"
	"dataStructures/arrays/arraysADT/binSearch"
	delete2 "dataStructures/arrays/arraysADT/delete"
	"dataStructures/arrays/arraysADT/display"
	"dataStructures/arrays/arraysADT/insert"
	"dataStructures/arrays/arraysADT/linSearch"
)

type Array struct {
	binSearch.Array
}

func (arr Array) Get(index int) int {
	if index >= 0 && index < arr.Length {
		return arr.A[index]
	}
	return -1
}

type ImplS1A struct{}

func (s1 ImplS1A) get(base [20]int, index int) int {
	var i int
	for _, v := range base {
		if v != 0 {
			i++
		}
	}

	arr := Array{binSearch.Array{linSearch.Array{delete2.Array{insert.Array{append2.Array{display.Array{base, 20, i}}}}}}}
	return arr.Get(index)
}
