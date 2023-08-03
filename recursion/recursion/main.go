package main

import (
	"fmt"
	"strconv"
	"strings"
)

var out string

func recursion(n int) string {
	if n > 0 {
		out += strconv.Itoa(n) + " "
		recursion(n - 1)
	}
	return strings.TrimSpace(out)
}

func main() {
	fmt.Println(recursion(3))
}
