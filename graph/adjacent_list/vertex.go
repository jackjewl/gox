package adjacent_list

import "gox/linear/list"

const (
	//顶点状态
	UnDiscovered = iota + 1
	Discovered
	Visited
)

type Vertex[VT any, ET any] struct {
	data      VT
	inDegree  int
	outDegree int
	status    int
	dTime     int
	fTime     int
	parent    int
	priority  int
	outEdges  *list.List[Edge[ET]]
	inEdges   *list.List[Edge[ET]]
}
