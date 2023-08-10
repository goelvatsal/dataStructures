package RecursionTailHead

import (
	"strconv"
	"strings"
)

var out string

type ImplS1R struct{}

func (s1 ImplS1R) recursionTailHead(n int) string {
	if n > 0 {
		out += strconv.Itoa(n) + " "
		s1.recursionTailHead(n - 1)
	}
	return strings.TrimSpace(out)
}
