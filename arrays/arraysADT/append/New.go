package append

import (
	//arrays "dataStructures/arrays/arraysADT/display"
	"dataStructures/arrays/arraysADT/display"
)

type MyArray struct {
	display.Array
}

func (arr MyArray) Append(x int) {
	arr.A[arr.Length] = x
	arr.Length++
}
