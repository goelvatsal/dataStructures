package append

type ArrayAppender interface {
	arrayAppender(init [20]int, new int) string
}
