package append

import (
	"dataStructures/arrays/arraysADT/display"
)

type Array struct {
	display.Array
}

func (arr *Array) Append(x int) {
	arr.A[arr.Length] = x
	arr.Length++
}

type ImplS1A struct{}

func (s1 ImplS1A) append(init [20]int, new int) string {
	var i int
	for _, v := range init {
		if v != 0 {
			i++
		}
	}

	arr := Array{display.Array{A: init, Size: 20, Length: i}}
	arr.Append(new)
	return arr.Display()
}
