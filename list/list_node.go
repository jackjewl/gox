package list

type ListNode struct {
	data interface{}
	pred *ListNode
	succ *ListNode
}

func (this *ListNode) Data() interface{} {
	return this.data
}

func (this *ListNode) Pred() *ListNode {
	return this.pred
}

func (this *ListNode) Succ() *ListNode {
	return this.succ
}

func (this *ListNode) InsertAsPred(value interface{}) {
	listNode := &ListNode{
		data: value,
		pred: this.pred,
		succ: this,
	}
	if this.pred != nil {
		this.pred.succ = listNode
	}

	this.pred = listNode
}

func (this *ListNode) InsertAsSucc(value interface{}) {
	listNode := &ListNode{
		data: value,
		pred: this,
		succ: this.succ,
	}
	if this.succ != nil {
		this.succ.pred = listNode
	}
	this.succ = listNode
}
