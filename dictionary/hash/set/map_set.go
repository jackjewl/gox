package set

type MapSet[T any] struct {
	mapper map[string]HashCoder
}
