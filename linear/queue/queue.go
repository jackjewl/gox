package queue

type Queue[T any] interface {
	Size() int
	Empty() bool
	EnQueue(T)
	DeQueue() T
	Front() T
}
