package difference

import "dataStructures/arrays/arraysADT/intersection"

type Array struct {
	intersection.Array
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

func (arr3 Array) difference(arr, arr2 [20]int) string {
	arrLen := lenCalc(arr)
	arr2Len := lenCalc(arr2)
	var i, j, k int

	for i < arrLen && j < arr2Len {
		if arr[i] < arr2[j] {
			arr3.Array.Array.Array.Array.A[k] = arr[i]
			k, i = k+1, i+1
		} else if arr2[j] < arr[i] {
			j++
		} else {
			i, j = i+1, j+1
		}
	}
	arr3.Array.Array.Array.Array.Length = k

	for ; i < arrLen; i++ {
		arr3.Array.Array.Array.Array.A[k] = arr[i]
		k++
	}
	return arr3.Array.Array.Array.Array.Display()
}

type ImplS1A struct{}

func (s1 ImplS1A) arraysDifference(arr, arr2 [20]int) string {
	var arr3 Array
	return arr3.difference(arr, arr2)
}
