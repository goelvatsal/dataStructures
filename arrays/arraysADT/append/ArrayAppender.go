package append

type Appender interface {
	append(base [20]int, valToAdd int) string
}
