package recursionNCR

type ImplS1R struct{}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return fact(n-1) * n
}

func (s1 ImplS1R) combinationFormula(n, r int) int {
	num := fact(n)
	den := fact(r) * fact(n-r)
	return num / den
}
