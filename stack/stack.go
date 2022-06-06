package stack

type Stack[T any] interface {
	Empty() bool
	Size() int
	Push(T)
	Pop() T
	Top() T
}
