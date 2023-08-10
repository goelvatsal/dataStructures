package main

import (
	"fmt"
	"strconv"
	"strings"
)

var out string

func recursionTree(n int) string {
	if n > 0 {
		out += strconv.Itoa(n) + " "
		recursionTree(n - 1)
		recursionTree(n - 1)
	}
	return strings.TrimSpace(out)
}

func main() {
	fmt.Println(recursionTree(3))
}
