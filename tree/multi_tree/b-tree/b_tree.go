package btree

type BTree struct {
	size  int
	order int
	root  *BTreeNode
	hot   *BTreeNode
}

func (this *BTree) SolveOverFlow(T *BTreeNode) {

}

func (this *BTree) SolveUnderFlow(T *BTreeNode) {

}

func NewBTree(order int) *BTree {
	btree := &BTree{
		size:  0,
		order: order,
		root:  NewBTreeNode(),
		hot:   nil,
	}
	return btree
}

func (this *BTree) Size() int {
	return this.size
}

func (this *BTree) Order() int {
	return this.order
}

func (this *BTree) Root() *BTreeNode {
	return this.root
}

func (this *BTree) Empty() bool {
	return this.Size() == 0
}

func (this *BTree) Search(target interface{}) *BTreeNode {

	return nil
}

func (this *BTree) Insert(value interface{}) *BTreeNode {
	return nil
}

func (this *BTree) Remove(value interface{}) *BTreeNode {
	return nil
}
