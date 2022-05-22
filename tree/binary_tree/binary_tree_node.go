package binarytree

const (
	Red int = iota + 1
	Black
)

type BinaryTreeNode[T any] struct {
	data                   T
	parent, lchild, rchild *BinaryTreeNode[T]
	height                 int
	npl                    int
	rbColor                int
}

func (this *BinaryTreeNode[T]) Size() int {

	return 0
}

func (this *BinaryTreeNode[T]) InsertAsLeftChild(data T) *BinaryTreeNode[T] {

	node := &BinaryTreeNode[T]{
		data:    data,
		parent:  this,
		lchild:  nil,
		rchild:  nil,
		height:  0,
		npl:     1,
		rbColor: Red,
	}

	this.lchild = node
	return node

}

func (this *BinaryTreeNode[T]) InsertAsRightChild(data T) *BinaryTreeNode[T] {
	node := &BinaryTreeNode[T]{
		data:    data,
		parent:  this,
		lchild:  nil,
		rchild:  nil,
		height:  0,
		npl:     1,
		rbColor: Red,
	}

	this.rchild = node
	return node

}

func (this *BinaryTreeNode[T]) Succ() *BinaryTreeNode[T] {

	s := this
	if this.HasRChild() {
		s = this.rchild
		for s.HasLChild() {
			s = s.lchild
		}
	} else {
		for s.IsRChild() {
			s = s.parent
		}
		s = s.parent
	}

	return s
}

func (this *BinaryTreeNode[T]) TraverseLevel() {

}

func (this *BinaryTreeNode[T]) TraversePre() {

}

func (this *BinaryTreeNode[T]) TraverseIn() {

}

func (this *BinaryTreeNode[T]) TraversePost() {

}

func (this *BinaryTreeNode[T]) Stature() int {

	if this == nil {
		return -1
	} else {
		return this.height
	}
}

func (this *BinaryTreeNode[T]) IsRoot() bool {

	if this.parent == nil {
		return true
	}
	return false
}
func (this *BinaryTreeNode[T]) IsLChild() bool {
	if this.parent.lchild == this {
		return true
	}
	return false
}
func (this *BinaryTreeNode[T]) IsRChild() bool {

	if this.parent.rchild == this {
		return true
	}
	return false
}
func (this *BinaryTreeNode[T]) HasParent() bool {

	return !this.IsRoot()
}
func (this *BinaryTreeNode[T]) HasLChild() bool {

	return this.lchild != nil
}
func (this *BinaryTreeNode[T]) HasRChild() bool {

	return this.rchild != nil
}
func (this *BinaryTreeNode[T]) HasChild() bool {

	return this.lchild != nil || this.rchild != nil
}
func (this *BinaryTreeNode[T]) HasBothChild() bool {
	return this.lchild != nil && this.rchild != nil
}
func (this *BinaryTreeNode[T]) IsLeaf() bool {
	return !this.HasChild()
}
func (this *BinaryTreeNode[T]) Sibling() *BinaryTreeNode[T] {

	if this.IsLChild() {
		return this.parent.rchild
	} else {
		return this.parent.lchild
	}
}
func (this *BinaryTreeNode[T]) Uncle() *BinaryTreeNode[T] {
	if this.parent.IsLChild() {
		return this.parent.rchild
	} else {
		return this.parent.lchild
	}
}
func (this *BinaryTreeNode[T]) FromParentTo() *BinaryTreeNode[T] {

	if this.IsRoot() {
		return nil
	} else {
		if this.IsRChild() {
			return this.parent.lchild
		} else {
			return this.parent.rchild
		}
	}

}
