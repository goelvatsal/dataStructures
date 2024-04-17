package delete

type ArrayDeleter interface {
	arrayDeleter(init [20]int, index int) string
}
