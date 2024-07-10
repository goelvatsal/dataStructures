package multipleMissingElements

import "dataStructures/arrays/arraysADT/missingElement"

type Array struct {
	Array missingElement.Array
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

func (arr3 Array) multipleMissingElements(arr [20]int) []int {
	l := lenCalc(arr)

	low := arr[0]
	n := arr[(l-1)/2]
	diff := low
	var opt []int

	for i := 0; i < n; i++ {
		if !(i >= 20) {
			if arr[i]-i != diff {
				for diff < arr[i]-i {
					opt = append(opt, i+diff)
					diff++
				}
			}
		}
	}
	return opt
}

type ImplS1A struct{}

func (s1 ImplS1A) arraysMultipleMissing(arr [20]int) []int {
	var arr3 Array
	return arr3.multipleMissingElements(arr)
}
