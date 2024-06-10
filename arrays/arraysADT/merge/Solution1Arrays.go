package merger

import (
	"dataStructures/arrays/arraysADT/negSort"
)

type Array struct {
	Array negSort.Array
}

func lenCalc(a [20]int) int {
	var i int
	for _, v := range a {
		if v != 0 {
			i++
		} else {
			break
		}
	}
	return i
}

func (arr3 Array) merge(arr, arr2 [20]int) string {
	var i, j, k int
	arrLen := lenCalc(arr)
	arr2Len := lenCalc(arr2)

	for i < arrLen && j < arr2Len {
		var val int
		if arr[i] < arr2[j] {
			val = arr[i]
			i++
		} else {
			val = arr2[j]
			j++
		}

		arr3.Array.A[k] = val
		k++
	}

	for ; i < arrLen; i++ {
		arr3.Array.A[k] = arr[i]
		k++
	}
	for ; j < arr2Len; j++ {
		arr3.Array.A[k] = arr2[j]
		k++
	}
	arr3.Array.Length = arrLen + arr2Len
	return arr3.Array.Display()
}

type ImplS1A struct{}

func (s1 ImplS1A) arraysMerger(arr, arr2 [20]int) string {
	//i := lenCalc(arr)
	//j := lenCalc(arr2)
	//var arr3 Array{negSort.Array{isSorted.Array{insertSort.Array{reverse.Array{sum.Array{maxMin.Array{set.Array{get.Array{binSearch.Array{linSearch.Array{delete2.Array{insert.Array{append2.Array{display.Array{[20]int{}, 20, i + j}}}}}}}}}}}}}}}
	var arr3 Array
	return arr3.merge(arr, arr2)
}
