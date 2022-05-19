package list

type List struct {
	size    int
	header  *ListNode
	trailer *ListNode
}

func NewList() *List {

	list := &List{
		size: 0,
		header: &ListNode{
			data: nil,
			pred: nil,
			succ: nil,
		},
		trailer: &ListNode{
			data: nil,
			pred: nil,
			succ: nil,
		},
	}
	list.header.succ = list.trailer
	list.trailer.pred = list.header
	return list
}

func (this *List) Size() int {
	return this.size
}

func (this *List) First() *ListNode {

	return this.header.succ
}

func (this *List) Last() ListNode {
	return *this.trailer.pred
}

func (this *List) Get(rank int) *ListNode {
	position := this.First()
	for i := 0; i < rank; i++ {
		position = position.succ
	}
	return position
}

func (this *List) FindSection(target interface{}, sectionNum int, end *ListNode, equal func(interface{}, interface{}) bool) *ListNode {

	for i := sectionNum; i > 0; i-- {
		end = end.pred
		if equal(end.data, target) {
			return end
		}
	}

	j
	return nil
}

func (this *List) Find(target interface{}, equal func(interface{}, interface{}) bool) *ListNode {

	return this.FindSection(target, this.size, this.trailer, equal)
}

func (this *List) InsertAsFirst(value interface{}) {

	listNode := ListNode{
		data: value,
		pred: this.header,
		succ: this.header.succ,
	}
	this.header.succ.pred = &listNode
	this.header.succ = &listNode

}

func (this *List) InsertAsLast(value interface{}) {

	listNode := ListNode{
		data: value,
		pred: this.header,
		succ: this.header.succ,
	}
	this.header.succ.pred = &listNode
	this.header.succ = &listNode

}

func (this *List) InsertBefroe(position *ListNode, value interface{}) {
	listNode := &ListNode{
		data: value,
		pred: position.pred,
		succ: position,
	}
	position.pred.succ = listNode
	position.pred = listNode
}

func (this *List) InsertAfter(position *ListNode, value interface{}) {
	listNode := &ListNode{
		data: value,
		pred: position,
		succ: position.succ,
	}
	position.succ.pred = listNode
	position.succ = listNode
}

func (this *List) CopyNodes(startPtr *ListNode, nodesNum int) {

	this = NewList()

	for i := 0; i < nodesNum; i++ {
		this.InsertAsLast(startPtr.data)
		startPtr = startPtr.succ
	}

}

func (this *List) CopyList(src *List) {
	this.CopyNodes(src.First(), src.size)
}

func (this *List) Remove(position *ListNode) interface{} {

	position.pred.succ = position.succ
	position.succ.pred = position.pred

}

func (this *List) Disordered() bool {

	return false
}

func (this *List) Sort(less func(interface{}, interface{}) bool) {

}

func (this *List) Search(less func(interface{}, interface{}) bool) int {

	return -1
}

func (this *List) Deduplicated(equal func(interface{}, interface{}) bool) int {
	return -1
}

func (this *List) Uniquify(equal func(interface{}, interface{}) bool) int {
	return -1
}

func (this *List) Traverse(visitor func(interface{}) bool) bool {

	return false
}
