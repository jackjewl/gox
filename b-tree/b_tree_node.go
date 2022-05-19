package btree

import (
	"gox/vector"
)

type BTreeNode struct {
	parent   *BTreeNode
	keys     *vector.Vector
	children *vector.Vector
}

func NewBTreeNode() *BTreeNode {
	node := &BTreeNode{
		parent:   nil,
		keys:     vector.NewVector(0),
		children: vector.NewVector(0),
	}
	node.children.Append(nil)
	return node
}
func NewBTreeNodeWithChildren() *BTreeNode {
	node := &BTreeNode{
		parent:   nil,
		keys:     vector.NewVector(1),
		children: vector.NewVector(2),
	}
	return node
}
