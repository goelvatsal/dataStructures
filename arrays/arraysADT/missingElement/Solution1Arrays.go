package missingElement

import (
	"dataStructures/arrays/arraysADT/difference"
)

type Array struct {
	Array difference.Array
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

func (arr2 Array) missingElement(arr [20]int) int {
	l := lenCalc(arr)

	low := arr[0]
	n := arr[(l-1)/2]
	diff := low - 0

	for i := 0; i < n; i++ {
		if arr[i]-i != diff {
			return i + diff
		}
	}
	return -1
}

type ImplS1A struct{}

func (s1 ImplS1A) arraysMissingElement(arr [20]int) int {
	var arr2 Array
	return arr2.missingElement(arr)
}
