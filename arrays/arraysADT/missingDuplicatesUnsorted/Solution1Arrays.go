package missingDuplicatesUnsorted

import "dataStructures/arrays/arraysADT/missingDuplicates"

type Array struct {
	Array missingDuplicates.Array
}

func lenCalc(arr [20]int) int {
	var n int
	for _, v := range arr {
		if v != 0 {
			n++
		}
	}
	return n
}

func (arr3 Array) missingDuplicatesUnsorted(arr [20]int) []int {
	l := lenCalc(arr)
	var fin []int

	for i := 0; i < l; i++ {
		count := 1
		for j := i + 1; j < l; j++ {
			if arr[i] == arr[j] {
				count++
			}
		}
		if count == 2 {
			fin = append(fin, arr[i])
		}
	}
	return fin
}

type ImplS1A struct{}

func (s1 ImplS1A) arraysMissingDuplicatesUnsorted(arr [20]int) []int {
	var arr3 Array
	return arr3.missingDuplicatesUnsorted(arr)
}
