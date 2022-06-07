package redblacktree

const (
	Red int = iota + 1
	Black
)

type Color int

type TreeNode struct {
	parent, lchild, rchild *TreeNode
	blackHeight            int
	color                  Color
}

//中序遍历的后继
func (this *TreeNode) Succ() *TreeNode {

	if this.rchild != nil {
		succ := this.rchild
		for succ != nil {
			succ = succ.lchild
		}
		return succ
	} else {
		succ := this
		for succ.parent.rchild == succ {
			succ = succ.parent
		}
		succ = succ.parent
		return succ
	}
}
