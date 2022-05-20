package list

type ListNode[T any] struct {
	data T
	pred *ListNode[T]
	succ *ListNode[T]
}

func (this *ListNode[T]) Data() T {
	return this.data
}

func (this *ListNode[T]) Pred() *ListNode[T] {
	return this.pred
}

func (this *ListNode[T]) Succ() *ListNode[T] {
	return this.succ
}

func (this *ListNode[T]) InsertAsPred(value T) {
	listNode := &ListNode[T]{
		data: value,
		pred: this.pred,
		succ: this,
	}
	if this.pred != nil {
		this.pred.succ = listNode
	}

	this.pred = listNode
}

func (this *ListNode[T]) InsertAsSucc(value T) {
	listNode := &ListNode[T]{
		data: value,
		pred: this,
		succ: this.succ,
	}
	if this.succ != nil {
		this.succ.pred = listNode
	}
	this.succ = listNode
}
