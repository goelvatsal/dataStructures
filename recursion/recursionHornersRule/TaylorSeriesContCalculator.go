package TaylorSeriesContinued

type TaylorSeriesHornerCalculator interface {
	recursionHornersRule(x, n, accumulated float64) float64
}
