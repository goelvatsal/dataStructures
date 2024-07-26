package set

import (
	append2 "dataStructures/arrays/arraysADT/append"
	"dataStructures/arrays/arraysADT/binSearch"
	delete2 "dataStructures/arrays/arraysADT/delete"
	"dataStructures/arrays/arraysADT/display"
	"dataStructures/arrays/arraysADT/get"
	"dataStructures/arrays/arraysADT/insert"
	"dataStructures/arrays/arraysADT/linSearch"
)

type Array struct {
	get.Array
}

func (arr *Array) Set(index int, newVal int) {
	if index >= 0 && index < arr.Size {
		if index == 0 {
			arr.A[index] = newVal
		} else {
			arr.A[index-1] = newVal
		}
	}
}

type ImplS1A struct{}

func (s1 ImplS1A) set(base [20]int, index int, newVal int) string {
	var i int
	for _, v := range base {
		if v != 0 {
			i++
		} else {
			break
		}
	}

	arr := Array{get.Array{binSearch.Array{linSearch.Array{delete2.Array{insert.Array{append2.Array{display.Array{base, 20, i}}}}}}}}
	arr.Set(index, newVal)
	return arr.Display()
}
