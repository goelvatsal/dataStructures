package get

type Getter interface {
	get(base [20]int, index int) int
}
