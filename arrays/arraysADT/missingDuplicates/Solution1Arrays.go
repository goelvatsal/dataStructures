package missingDuplicates

import "dataStructures/arrays/arraysADT/multipleMissingElements"

type Array struct {
	Array multipleMissingElements.Array
}

func lenCalc(arr [20]int) int {
	var i int
	for _, v := range arr {
		if v != 0 {
			i++
		} else {
			break
		}
	}
	return i
}

func (arr3 Array) missingDuplicates(arr [20]int) []int {
	l := lenCalc(arr)
	var opt []int
	lastDup := 0
	for i := 0; i < l; i++ {
		if arr[i] == arr[i+1] && lastDup != arr[i] {
			opt = append(opt, arr[i])
			lastDup = arr[i]
		}
	}
	return opt
}

type ImplS1A struct{}

func (s1 ImplS1A) arraysMissingDuplicates(arr [20]int) []int {
	var arr3 Array
	return arr3.missingDuplicates(arr)
}
