package stack

import "gox/linear/vector"

type SequenceStack[T any] struct {
	Elems *vector.Vector[T]
}

func NewSequenceStack[T any]() *SequenceStack[T] {
	stack := &SequenceStack[T]{
		Elems: vector.NewVector[T](0),
	}
	return stack
}

func (this *SequenceStack[T]) Empty() bool {
	return this.Elems.Size() == 0
}

func (this *SequenceStack[T]) Size() int {
	return this.Elems.Size()

}
func (this *SequenceStack[T]) Push(value T) {
	this.Elems.Append(value)

}
func (this *SequenceStack[T]) Pop() T {
	removedValue := this.Elems.Get(this.Elems.Size() - 1)
	this.Elems.Remove(this.Elems.Size() - 1)
	return removedValue

}
func (this *SequenceStack[T]) Top() T {
	return this.Elems.Get(this.Elems.Size() - 1)
}
