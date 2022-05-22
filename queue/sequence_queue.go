package queue

import "gox/vector"

type SequenceQueue[T any] struct {
	Elems *vector.Vector[T]
}

func NewSequenceQueue[T any]() *SequenceQueue[T] {
	queue := &SequenceQueue[T]{
		Elems: vector.NewVector[T](0),
	}
	return queue
}
func (this *SequenceQueue[T]) Size() int {
	return this.Elems.Size()
}

func (this *SequenceQueue[T]) Empty() bool {
	return this.Size() == 0
}

func (this *SequenceQueue[T]) EnQueue(value T) {
	this.Elems.Append(value)
}

func (this *SequenceQueue[T]) DeQueue() T {
	removedValue := this.Elems.Get(0)
	this.Elems.Remove(0)
	return removedValue
}
func (this *SequenceQueue[T]) Front() T {
	return this.Elems.Get(0)
}
