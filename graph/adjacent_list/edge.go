package adjacent_list

const (
	//边状态
	Undiscovered = iota + 1
	Tree
	Cross
	Forward
	Backward
)

type Edge[ET any] struct {
	data            ET
	weight          int
	status          int
	edgeHeaderIndex int
}
