package Arrays2D

import "fmt"

type ImplS1A struct{}

func (s1 ImplS1A) arrays2D() string {
	arr := [3][4]int{
		{1, 2, 3, 4},
		{2, 4, 6, 8},
		{1, 3, 5, 7},
	}

	var out string
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			out += fmt.Sprintf("%d ", arr[i][j])
		}
		out += fmt.Sprintln()
	}
	return out
}
