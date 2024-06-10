package intersection

import "dataStructures/arrays/arraysADT/union"

type Array struct {
	union.Array
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

func (arr3 Array) intersection(arr, arr2 [20]int) string {
	var i, j, k int
	arrLen := lenCalc(arr)
	arr2Len := lenCalc(arr2)

	for i < arrLen && j < arr2Len {
		if arr[i] < arr2[j] {
			i++
		} else if arr2[j] < arr[i] {
			j++
		} else if arr[i] == arr2[j] {
			arr3.Array.Array.Array.A[k] = arr[i]
			k, j, i = k+1, j+1, i+1
		}
	}
	arr3.Array.Array.Array.Length = k
	return arr3.Array.Array.Array.Display()
}

type ImplS1A struct{}

func (s1 ImplS1A) arraysIntersection(arr, arr2 [20]int) string {
	var arr3 Array
	return arr3.intersection(arr, arr2)
}
