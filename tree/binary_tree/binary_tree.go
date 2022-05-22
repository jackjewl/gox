package binarytree

import (
	"gox/queue"
	"gox/stack"
	"gox/utils"
)

type BinaryTree[T any] struct {
	size int
	root *BinaryTreeNode[T]
}

func NewBinaryTree[T any]() *BinaryTree[T] {

	btree := &BinaryTree[T]{
		size: 0,
		root: nil,
	}
	return btree

}

func (this *BinaryTree[T]) UpdateHeight(ptr *BinaryTreeNode[T]) {

	ptr.height = utils.Max[int](ptr.rchild.height, ptr.lchild.height)

}

func (this *BinaryTree[T]) UpdateHeightAbove(ptr *BinaryTreeNode[T]) {
	for ptr != nil {
		this.UpdateHeight(ptr)
		ptr = ptr.parent
	}

}

func (this *BinaryTree[T]) Size(ptr *BinaryTreeNode[T]) int {
	panic("impl me")

}

func (this *BinaryTree[T]) Empty(ptr *BinaryTreeNode[T]) bool {
	panic("impl me")

}

func (this *BinaryTree[T]) Root(ptr *BinaryTreeNode[T]) *BinaryTreeNode[T] {
	panic("impl me")
}

func (this *BinaryTree[T]) InsertAsRoot(data T) *BinaryTreeNode[T] {
	node := &BinaryTreeNode[T]{
		data:    data,
		parent:  nil,
		lchild:  nil,
		rchild:  nil,
		rbColor: Red,
		height:  0,
		npl:     1,
	}
	this.root = node
	this.size = 1
	return this.root

}

func (this *BinaryTree[T]) InsertAsLChild(ptr *BinaryTreeNode[T], data T) *BinaryTreeNode[T] {
	ptr.InsertAsLeftChild(data)
	this.size++
	this.UpdateHeightAbove(ptr)
	return ptr.lchild
}

func (this *BinaryTree[T]) InsertAsRChild(ptr *BinaryTreeNode[T], data T) *BinaryTreeNode[T] {
	ptr.InsertAsRightChild(data)
	this.size++
	this.UpdateHeightAbove(ptr)
	return ptr.rchild
}

func (this *BinaryTree[T]) AttachAsLSubTree(ptr *BinaryTreeNode[T], subTree *BinaryTree[T]) *BinaryTreeNode[T] {
	ptr.lchild = subTree.root
	ptr.lchild.parent = ptr
	this.size += subTree.size
	this.UpdateHeightAbove(ptr)
	subTree = nil
	return ptr
}

func (this *BinaryTree[T]) Remove(ptr *BinaryTreeNode[T]) int {
	if ptr.IsLChild() {
		ptr.parent.lchild = nil
	} else {
		ptr.parent.rchild = nil
	}

	this.UpdateHeightAbove(ptr.parent)
	removedNUm := this.RemoveAt(ptr)
	this.size -= removedNUm
	return removedNUm

}

func (this *BinaryTree[T]) RemoveAt(ptr *BinaryTreeNode[T]) int {
	if ptr == nil {
		return 0
	}
	counter := 0
	counter = 1 + this.RemoveAt(ptr.lchild) + this.RemoveAt(ptr.rchild)
	ptr = nil
	return counter

}

func (this *BinaryTree[T]) AttachAsRSubTree(ptr *BinaryTreeNode[T], subTree *BinaryTree[T]) *BinaryTreeNode[T] {
	ptr.rchild = subTree.root
	ptr.rchild.parent = ptr
	this.size += subTree.size
	this.UpdateHeightAbove(ptr)
	subTree = nil
	return ptr

}

func (this *BinaryTree[T]) Seced(ptr *BinaryTreeNode[T]) *BinaryTree[T] {

	if ptr.IsLChild() {
		ptr.parent.lchild = nil
	} else {
		ptr.parent.rchild = nil
	}
	this.UpdateHeightAbove(ptr.parent)

	ptr.parent = nil
	this.size -= ptr.Size()
	btree := &BinaryTree[T]{
		size: ptr.Size(),
		root: ptr,
	}
	return btree
}

func (this *BinaryTree[T]) TraverseLevel(visit func(T) bool) {

	q := queue.NewLinkQueue[*BinaryTreeNode[T]]()
	q.EnQueue(this.root)
	for !q.Empty() {
		h := q.DeQueue()
		visit(h.data)
		if h.HasLChild() {
			q.EnQueue(h.lchild)
		}
		if h.HasRChild() {
			q.EnQueue(h.rchild)
		}
	}

}
func (this *BinaryTree[T]) TraversePre(visit func(T) bool) {
	s := stack.NewLinkedStack[*BinaryTreeNode[T]]()
	ptr := this.root
	for true {
		this.VisitAlongLeftBranch(ptr, visit, s)
		if s.Empty() {
			break
		}
		ptr = s.Pop()
	}
}

func (this *BinaryTree[T]) VisitAlongLeftBranch(ptr *BinaryTreeNode[T], visit func(T) bool, s stack.Stack[*BinaryTreeNode[T]]) {
	for ptr != nil {
		visit(ptr.data)
		s.Push(ptr.rchild)
		ptr = ptr.lchild
	}
}
func (this *BinaryTree[T]) TraverseIn(visit func(T) bool) {
	s := stack.NewLinkedStack[*BinaryTreeNode[T]]()
	ptr := this.root
	for true {
		this.GoAlongLeftBranch(ptr, s)
		if s.Empty() {
			break
		}
		ptr = s.Pop()
		visit(ptr.data)
		ptr = ptr.rchild
	}

}

func (this *BinaryTree[T]) GoAlongLeftBranch(ptr *BinaryTreeNode[T], s stack.Stack[*BinaryTreeNode[T]]) {
	for ptr != nil {
		s.Push(ptr)
		ptr = ptr.lchild
	}
}
func (this *BinaryTree[T]) TraversePost(visit func(T) bool) {
	ptr := this.root
	s := stack.NewLinkedStack[*BinaryTreeNode[T]]()
	s.Push(ptr)
	for !s.Empty() {
		if s.Top() != ptr.parent {
			this.GotoHLVFL(s)
		}
		ptr = s.Pop()
		visit(ptr.data)
	}

}

func (this *BinaryTree[T]) GotoHLVFL(s stack.Stack[*BinaryTreeNode[T]]) {
	var ptr *BinaryTreeNode[T] = s.Top()
	for ptr != nil {
		ptr = s.Top()
		if ptr.HasLChild() {
			if ptr.HasRChild() {
				s.Push(ptr.rchild)
			}
			s.Push(ptr.lchild)
		} else {
			s.Push(ptr.rchild)
		}
	}
	s.Pop()
}
