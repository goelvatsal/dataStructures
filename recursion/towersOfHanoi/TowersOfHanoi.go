package TowersOfHanoi

type TOHCalculator interface {
	towersOfHanoi(n, a, b, c int) string
}
