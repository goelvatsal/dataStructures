package TowersOfHanoi

import "fmt"

type ImplS1R struct{}

var out string

func (s1 ImplS1R) towersOfHanoi(n, a, b, c int) string {
	if n > 0 {
		s1.towersOfHanoi(n-1, a, c, b)
		out += fmt.Sprintf("(%d, %d)\n", a, c)
		s1.towersOfHanoi(n-1, b, a, c)
	}
	return out
}
