package set

type Setter interface {
	set(base [20]int, index int, newVal int) string
}
