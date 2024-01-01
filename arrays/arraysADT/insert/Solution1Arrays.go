package insert

import (
	//arrays "dataStructures/arrays/arraysADT/display"
	append2 "dataStructures/arrays/arraysADT/append"
	"dataStructures/arrays/arraysADT/display"
	"fmt"
)

type Array struct {
	append2.Array
}

type ImplS1A struct{}

func (arr *Array) Insert(index int, x int) {
	if index >= 0 {
		for i := arr.Length; i > index; i-- {
			arr.A[i] = arr.A[i-1]
		}
		arr.A[index] = x
		arr.Length++
	} else {
		fmt.Println("Out of range!")
	}
}

func (s1 ImplS1A) Insert(init [20]int, index int, x int) string {
	var i int
	for _, v := range init {
		if v != 0 {
			i++
		}
	}

	arr := Array{append2.Array{display.Array{A: init, Size: 20, Length: i}}}
	arr.Insert(index, x)
	return arr.Display()

}
