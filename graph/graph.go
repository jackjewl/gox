package graph

type Graph[VT any, ET any] interface {
	//对边的操作接口
	EdgesNum() int
	EdgeExist(v, u int) bool
	EdgeInsert(v, u int)
	EdgeRemove(v, u int)
	EdgeStatus(v, u int) int
	EdgeData(v, u int) T
	Weight(v, u int) int
	//对顶点的操作接口
	VertexNum() int
	VertexInsert(vertexData T)
	VertexRemove(rank int)
	VertexInDegree(rank int) int
	VertexOutDegree(rank int) int
	VertexFirstNeighbor(rank int) int
	VertexNextNeighbor(rank, currentNeighbor int) int
	VertexStatus(rank int) int
	VertexDTime(rank int) int
	VertexFTime(rank int) int
	VertexParent(rank int) int
	VertexPriority(rank int) int
	//将所有边和顶点的状态复位，但是顶点和边的数据不变
	Reset()
}

//各种图的算法

func BFS[VT any, ET any](graph *Graph[VT, ET]) {

}

func BFSComponent[VT any, ET any](graph *Graph[VT, ET], statVertex int) {

}

func DFS[VT any, ET any](graph *Graph[VT, ET]) {

}

func DFSComponent[VT any, ET any](graph *Graph[VT, ET], statVertex int) {

}

func TopologicSort[VT any, ET any](graph *Graph[VT, ET], statVertex int) (sortSequence []int, err error) {

	return nil, nil

}

func BCC[VT any, ET any](graph *Graph[VT, ET]) {

}

func BCCComponent[VT any, ET any](graph *Graph[VT, ET], statVertex int) {

}

func PFS[VT any, ET any](graph *Graph[VT, ET], statVertex int) {

}

func Prim[VT any, ET any](graph *Graph[VT, ET], statVertex int) {

}

func Dijkstra[VT any, ET any](graph *Graph[VT, ET], statVertex int) {

}
