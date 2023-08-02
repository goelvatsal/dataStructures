package main

import (
	"unsafe"
)

type Rectangle struct {
	length  int
	breadth int
}

func pointerPractice() []int {
	var p1 = new(int)
	var p2 = new(rune)
	var p3 = new(float64)
	var p4 = new(Rectangle)

	result := append([]int{}, int(unsafe.Sizeof(p1)))
	result = append(result, int(unsafe.Sizeof(p2)))
	result = append(result, int(unsafe.Sizeof(p3)))
	result = append(result, int(unsafe.Sizeof(p4)))

	//fmt.Println(unsafe.Sizeof(p1))
	//fmt.Println(unsafe.Sizeof(p2))
	//fmt.Println(unsafe.Sizeof(p3))
	//fmt.Println(unsafe.Sizeof(p4))

	return result
}
