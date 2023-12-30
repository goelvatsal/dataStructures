package display

import (
	"fmt"
)

type Array struct {
	A      [20]int
	Size   int
	Length int
}

func (arr Array) Display() string {
	var out string

	out += fmt.Sprintf("Elements are: ")
	fmt.Printf("Elements are: ")
	for i := 0; i < arr.Length; i++ {
		out += fmt.Sprintf("%d ", arr.A[i])
		fmt.Printf("%d ", arr.A[i])
	}
	out += fmt.Sprintln()
	fmt.Println()
	return out
}

type ImplS1A struct{}

func (s1 ImplS1A) arrayPrinter() string {
	arr := Array{[20]int{2, 3, 4, 5, 6}, 20, 5}
	out := arr.Display()
	return out
}
