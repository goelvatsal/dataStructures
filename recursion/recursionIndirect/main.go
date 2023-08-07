package main

import (
	"fmt"
)

func a(n int, in *string) {
	if n > 0 {
		*in += fmt.Sprintf("%d ", n)
		b(n-1, in)
	}
}

func b(n int, in *string) {
	if n > 1 {
		*in += fmt.Sprintf("%d ", n)
		a(n/2, in)
	}
}

func main() {
	var pointer string
	a(20, &pointer)

	fmt.Println(pointer)

}
