package stack

import "gox/linear/list"

type LinkedStack[T any] struct {
	Elems *list.List[T]
}

func NewLinkedStack[T any]() *LinkedStack[T] {
	stack := &LinkedStack[T]{
		Elems: list.NewList[T](),
	}
	return stack
}

func (this *LinkedStack[T]) Empty() bool {
	return this.Elems.Empty()
}

func (this *LinkedStack[T]) Size() int {
	return this.Elems.Size()

}

func (this *LinkedStack[T]) Push(value T) {
	this.Elems.InsertAsLast(value)
}

func (this *LinkedStack[T]) Pop() T {
	removedValue := this.Elems.Last().Data()
	this.Elems.Remove(this.Elems.Last())
	return removedValue

}

func (this *LinkedStack[T]) Top() T {
	return this.Elems.Last().Data()
}
