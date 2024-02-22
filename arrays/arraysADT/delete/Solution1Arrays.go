package delete

import (
	append2 "dataStructures/arrays/arraysADT/append"
	"dataStructures/arrays/arraysADT/display"
	"dataStructures/arrays/arraysADT/insert"
)

type Array struct {
	insert.Array
}

func (arr *Array) Delete(index int) {
	if index >= 0 {
		for i := index; i < arr.Length-1; i++ {
			arr.A[i] = arr.A[i+1]
		}
		arr.Length--
		arr.A[arr.Length] = 0
	}
}

type Impl1SA struct{}

func (s1 Impl1SA) arrayDeleter(init [20]int, index int) string {
	var i int
	for _, v := range init {
		if v != 0 {
			i++
		}
	}

	arr := Array{insert.Array{append2.Array{display.Array{init, 20, i}}}}
	arr.Delete(index)
	return arr.Display()
}
