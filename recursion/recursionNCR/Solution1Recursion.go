package recursionNCR

type ImplS1R struct{}

func (s1 ImplS1R) fact(n int) int {
	if n == 0 {
		return 1
	}
	return s1.fact(n-1) * n
}

func (s1 ImplS1R) combinationFormula(n, r int) int {
	num := s1.fact(n)
	den := s1.fact(r) * s1.fact(n-r)
	return num / den
}
