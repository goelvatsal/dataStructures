package reverse

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
	"dataStructures/arrays/arraysADT/sum"
)

type Array struct {
	sum.Array
}

func (arr *Array) Reverse() string {
	B := arr.A

	for i, j := arr.Length-1, 0; i >= 0; i, j = i-1, j+1 {
		B[j] = arr.A[i]
	}
	for i := 0; i < arr.Length; i++ {
		arr.A[i] = B[i]
	}
	return arr.Display()
}

func (arr *Array) Reverse2() string {
	for i, j := 0, arr.Length-1; i < j; i, j = i+1, j-1 {
		arr.A[i], arr.A[j] = arr.A[j], arr.A[i]
	}
	return arr.Display()
}

type ImplS1A struct{}

func (s1 ImplS1A) arrayReverser(base [20]int) (string, string) {
	var i int
	for _, v := range base {
		if v != 0 {
			i++
		} else {
			break
		}
	}

	arr1 := Array{sum.Array{maxMin.Array{set.Array{get.Array{binSearch.Array{linSearch.Array{delete2.Array{insert.Array{append2.Array{display.Array{base, 20, i}}}}}}}}}}}
	arr2 := Array{sum.Array{maxMin.Array{set.Array{get.Array{binSearch.Array{linSearch.Array{delete2.Array{insert.Array{append2.Array{display.Array{base, 20, i}}}}}}}}}}}
	return arr1.Reverse(), arr2.Reverse2()
}
