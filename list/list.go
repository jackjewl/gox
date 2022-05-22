package list

import (
	"fmt"
	"log"
	"strconv"
)

type List[T any] struct {
	size    int
	header  *ListNode[T]
	trailer *ListNode[T]
}

func (this *List[T]) String() string {
	s := "size:" + strconv.Itoa(this.size) + "\telems: "
	for p := this.First(); p != this.trailer; p = p.succ {
		s += fmt.Sprintf("%v\t", p.data)
	}
	return s
}

func NewList[T any]() *List[T] {

	list := &List[T]{
		size: 0,
		header: &ListNode[T]{
			pred: nil,
			succ: nil,
		},
		trailer: &ListNode[T]{
			pred: nil,
			succ: nil,
		},
	}
	list.header.succ = list.trailer
	list.trailer.pred = list.header
	return list
}
func (this *List[T]) Clear() {
	*this = *NewList[T]()
}
func (this *List[T]) Size() int {
	return this.size
}

func (this *List[T]) Empty() bool {
	return this.size == 0
}

func (this *List[T]) Header() *ListNode[T] {
	return this.header
}

func (this *List[T]) Trailer() *ListNode[T] {
	return this.trailer
}

func (this *List[T]) First() *ListNode[T] {

	return this.header.succ
}

func (this *List[T]) Last() *ListNode[T] {
	return this.trailer.pred
}

func (this *List[T]) Get(rank int) *ListNode[T] {
	position := this.First()
	for i := 0; i < rank; i++ {
		position = position.succ
	}
	return position
}

func (this *List[T]) FindSection(target T, sectionNum int, end *ListNode[T], equal func(T, T) bool) *ListNode[T] {
	end = end.pred
	for i := sectionNum; i > 0; i-- {
		if equal(end.data, target) {
			return end
		}
		end = end.pred
	}
	return nil
}

func (this *List[T]) Find(target T, equal func(T, T) bool) *ListNode[T] {
	return this.FindSection(target, this.size, this.trailer, equal)
}

func (this *List[T]) InsertAsFirst(value T) {

	listNode := &ListNode[T]{
		data: value,
		pred: this.header,
		succ: this.header.succ,
	}
	this.header.succ.pred = listNode
	this.header.succ = listNode

	this.size++

}

func (this *List[T]) InsertAsLast(value T) {
	listNode := &ListNode[T]{
		data: value,
		pred: this.trailer.pred,
		succ: this.trailer,
	}
	this.trailer.pred.succ = listNode
	this.trailer.pred = listNode

	this.size++

}

func (this *List[T]) InsertBefore(position *ListNode[T], value T) {
	listNode := &ListNode[T]{
		data: value,
		pred: position.pred,
		succ: position,
	}
	position.pred.succ = listNode
	position.pred = listNode
	this.size++
}

func (this *List[T]) InsertAfter(position *ListNode[T], value T) {
	listNode := &ListNode[T]{
		data: value,
		pred: position,
		succ: position.succ,
	}
	position.succ.pred = listNode
	position.succ = listNode
	this.size++
}

func (this *List[T]) CopyNodes(startPtr *ListNode[T], nodesNum int) {

	for i := 0; i < nodesNum; i++ {
		this.InsertAsLast(startPtr.data)
		startPtr = startPtr.succ
	}
	this.size = nodesNum

}

func (this *List[T]) CopyNodesSection(src *List[T], startNodeIndex int, nodesNum int) {
	startPtr := src.Get(startNodeIndex)
	this.CopyNodes(startPtr, nodesNum)
	this.size = nodesNum
}

func (this *List[T]) CopyList(src *List[T]) {
	this.CopyNodes(src.First(), src.size)
	this.size = src.size
}

func (this *List[T]) Remove(position *ListNode[T]) T {

	removedValue := position.data
	position.pred.succ = position.succ
	position.succ.pred = position.pred
	this.size--
	return removedValue
}

func (this *List[T]) Disordered(less func(T, T) bool) bool {
	if this.Size() < 2 {
		return true
	}

	counter := 0
	ptr := this.First().Succ()
	for ptr != this.trailer {
		if less(ptr.data, ptr.pred.data) {
			counter++
		}
		ptr = ptr.succ
	}

	return counter == 0
}

func (this *List[T]) Deduplicated(equal func(T, T) bool) int {
	if this.Size() < 2 {
		return 0
	}
	oldSize := this.Size()
	ptr := this.First()
	rank := 0

	for ptr != this.trailer {
		if findResult := this.FindSection(ptr.data, rank, ptr, equal); findResult != nil {
			this.Remove(findResult)
		} else {
			rank++
		}
		ptr = ptr.succ
	}
	return oldSize - this.size
}

func (this *List[T]) Uniquify(equal func(T, T) bool) int {

	if this.Size() < 2 {
		return 0
	}
	oldSize := this.Size()
	pre, next := this.First(), this.First().Succ()
	for next != this.trailer {
		if equal(pre.data, next.data) {
			next = next.succ
			this.Remove(next.pred)
		} else {
			pre = next
			next = next.succ
		}
	}
	return oldSize - this.size
}

func (this *List[T]) Traverse(visitor func(T) bool) bool {

	ptr := this.First()
	if ptr == nil {
		return true
	}

	for ptr != this.trailer {
		if !visitor(ptr.Data()) {
			return false
		}
		ptr = ptr.succ
	}
	return true
}

//有序链表查找
func (this *List[T]) Search(target T, n int, endPosition *ListNode[T], compare func(T, T) int) *ListNode[T] {
	endPosition = endPosition.pred
	for n >= 0 {
		if compare(target, endPosition.data) >= 0 {
			return endPosition
		}
		endPosition = endPosition.pred
		n--
	}
	return endPosition
}

func (this *List[T]) InsertionSort(startPosition *ListNode[T], num int, compare func(T, T) int) {
	if num < 2 {
		return
	}

	for i := 0; i < num; i++ {
		this.InsertAfter(this.Search(startPosition.data, i, startPosition, compare), startPosition.data)
		startPosition = startPosition.succ
		this.Remove(startPosition.pred)
	}
}

func (this *List[T]) SelectionSort(start *ListNode[T], num int, compare func(T, T) int) {
	if num < 2 {
		return
	}
	head, tail := start, start.succ
	for i := 0; i < num; i++ {
		tail = tail.succ
	}

	for num > 1 {
		max := this.SelectMax(head, num, compare)
		this.InsertBefore(tail, max.data)
		this.Remove(max)
		tail = tail.pred
		num--
	}

}

func (this *List[T]) SelectMax(startPosition *ListNode[T], num int, compare func(T, T) int) *ListNode[T] {
	max := startPosition
	startPosition = startPosition.succ
	for num > 1 {
		if compare(startPosition.data, max.data) > 0 {
			max = startPosition
		}
		startPosition = startPosition.succ
		num--
	}
	return max
}

func (this *List[T]) MergeSort(compare func(T, T) int) {
	this.MergeSortR(0, this.Size()-1, compare)
}
func (this *List[T]) MergeSortR(start, end int, compare func(T, T) int) {
	log.Println("rc")
	if end-start < 2 {
		return
	}

	mid := (start + end) >> 1

	this.MergeSortR(start, mid, compare)
	this.MergeSortR(mid+1, end, compare)
	log.Println(this)
	this.Merge(this.Get(start), this.Get(mid), this.Get(end), compare)
}

//TODO
func (this *List[T]) Merge(start, mid, end *ListNode[T], compare func(T, T) int) {

	leftPtr, midPtr, ptr := start, mid.succ, start.pred
	//100,2,32,400,55
	for leftPtr != mid.succ && midPtr != end.succ {
		//log.Println(leftPtr.data, midPtr.data)
		if compare(leftPtr.data, midPtr.data) > 0 {
			ptr.succ = midPtr
			midPtr.pred = ptr
			midPtr = midPtr.succ
		} else {
			ptr.succ = leftPtr
			leftPtr.pred = ptr
			leftPtr = leftPtr.succ
		}
		ptr = ptr.succ
	}
	if leftPtr != mid.succ {
		ptr.succ = leftPtr
		leftPtr.pred = ptr
	}
	if midPtr != end.succ {
		ptr.succ = leftPtr
		midPtr.pred = ptr
	}
}
