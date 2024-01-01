package insert

type Inserter interface {
	Insert(base [20]int, index int, x int) string
}
