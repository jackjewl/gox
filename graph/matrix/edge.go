package matrix

type Edge[T any] struct {
	data   T
	weight int
	status int
}
