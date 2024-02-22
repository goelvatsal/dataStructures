package insert

type ArrayInserter interface {
	arrayInserter(init [20]int, index int, x int) string
}
