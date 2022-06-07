package queue

import "gox/linear/list"

//linked queue
type LinkQueue[T any] struct {
	Elems *list.List[T]
}

//constuct function
func NewLinkQueue[T any]() *LinkQueue[T] {

	queue := &LinkQueue[T]{
		Elems: list.NewList[T](),
	}
	return queue
}

func (this *LinkQueue[T]) Size() int {
	return this.Elems.Size()
}

func (this *LinkQueue[T]) Empty() bool {

	return this.Size() == 0
}

func (this *LinkQueue[T]) EnQueue(values T) {
	this.Elems.InsertAsLast(values)
}

func (this *LinkQueue[T]) DeQueue() T {
	removedValue := this.Elems.First().Data()
	this.Elems.Remove(this.Elems.First())
	return removedValue
}

func (this *LinkQueue[T]) Front() T {
	return this.Elems.First().Data()
}
