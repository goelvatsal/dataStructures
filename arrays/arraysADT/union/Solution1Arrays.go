package union

import merger "dataStructures/arrays/arraysADT/merge"

type Array struct {
	merger.Array
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

func (arr3 Array) union(arr, arr2 [20]int) string {
	var i, j, k int
	arrLen := lenCalc(arr)
	arr2Len := lenCalc(arr2)

	for i < arrLen && j < arr2Len {
		var val int
		if arr[i] < arr2[j] {
			val = arr[i]
			i++
		} else if arr2[j] < arr[i] {
			val = arr2[j]
			j++
		} else {
			val = arr[i]
			i++
			j++
		}
		arr3.Array.Array.A[k] = val
		k++
	}
	for ; i < arrLen; i++ {
		arr3.Array.Array.A[k] = arr[i]
		k++
	}
	for ; j < arr2Len; j++ {
		arr3.Array.Array.A[k] = arr2[j]
		k++
	}
	arr3.Array.Array.Length = k
	return arr3.Array.Array.Display()
}

type ImplS1A struct{}

func (s1 ImplS1A) arraysUnion(arr, arr2 [20]int) string {
	var arr3 Array
	return arr3.union(arr, arr2)
}
