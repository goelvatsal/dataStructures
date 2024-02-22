package maxMin

import (
	append2 "dataStructures/arrays/arraysADT/append"
	"dataStructures/arrays/arraysADT/binSearch"
	delete2 "dataStructures/arrays/arraysADT/delete"
	"dataStructures/arrays/arraysADT/display"
	"dataStructures/arrays/arraysADT/get"
	"dataStructures/arrays/arraysADT/insert"
	"dataStructures/arrays/arraysADT/linSearch"
	"dataStructures/arrays/arraysADT/set"
)

type Array struct {
	set.Array
}

func (arr Array) Max() int {
	biggest := arr.A[0]
	for i := 0; i < arr.Length; i++ {
		if arr.A[i] > biggest {
			biggest = arr.A[i]
		}
	}
	return biggest
}

func (arr Array) Min() int {
	smallest := arr.A[0]
	for i := 0; i < arr.Length; i++ {
		if arr.A[i] < smallest {
			smallest = arr.A[i]
		}
	}
	return smallest
}

type ImplS1A struct{}

func (s1 ImplS1A) maxMin(base [20]int) (int, int) {
	var i int
	for _, v := range base {
		if v != 0 {
			i++
		} else {
			break
		}
	}

	arr := Array{set.Array{get.Array{binSearch.Array{linSearch.Array{delete2.Array{insert.Array{append2.Array{display.Array{base, 20, i}}}}}}}}}
	return arr.Max(), arr.Min()
}

//func (s1 ImplS1A) min(base [20]int) int {
//	var i int
//	for _, v := range base {
//		if v != 0 {
//			i++
//		} else {
//			break
//		}
//	}
//
//	arr := Array{set.Array{get.Array{binSearch.Array{linSearch.Array{delete2.Array{insert.Array{append2.Array{display.Array{base, 20, i}}}}}}}}}
//	return arr.Min()
//}
