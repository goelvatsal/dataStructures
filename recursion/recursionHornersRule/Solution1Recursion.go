package TaylorSeriesContinued

type S1R struct{}

func (s1 S1R) recursionHornersRule(x, n, accumulated float64) float64 {
	if n == 0 {
		return accumulated
	}

	newAccumulated := 1 + x*accumulated/n
	return s1.recursionHornersRule(x, n-1, newAccumulated)
}
