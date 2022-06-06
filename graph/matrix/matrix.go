package matrix

import "gox/vector"

type GraphMatrix[VT any, ET any] struct {
	vertexNum int
	edgeNum   int
	vertexes  vector.Vector[Vertex[VT]]
	edges     vector.Vector[vector.Vector[Edge[ET]]]
}

func (g GraphMatrix) EdgesNum() int {
	//TODO implement me
	panic("implement me")
}

func (g GraphMatrix) EdgeExist(v, u int) bool {
	//TODO implement me
	panic("implement me")
}

func (g GraphMatrix) EdgeInsert(v, u int) {
	//TODO implement me
	panic("implement me")
}

func (g GraphMatrix) EdgeRemove(v, u int) {
	//TODO implement me
	panic("implement me")
}

func (g GraphMatrix) EdgeStatus(v, u int) int {
	//TODO implement me
	panic("implement me")
}

func (g GraphMatrix) EdgeData(v, u int) interface{} {
	//TODO implement me
	panic("implement me")
}

func (g GraphMatrix) Weight(v, u int) int {
	//TODO implement me
	panic("implement me")
}

func (g GraphMatrix) VertexNum() int {
	//TODO implement me
	panic("implement me")
}

func (g GraphMatrix) VertexInsert(vertexData interface{}) {
	//TODO implement me
	panic("implement me")
}

func (g GraphMatrix) VertexRemove(rank int) {
	//TODO implement me
	panic("implement me")
}

func (g GraphMatrix) VertexInDegree(rank int) int {
	//TODO implement me
	panic("implement me")
}

func (g GraphMatrix) VertexOutDegree(rank int) int {
	//TODO implement me
	panic("implement me")
}

func (g GraphMatrix) VertexFirstNeighbor(rank int) int {
	//TODO implement me
	panic("implement me")
}

func (g GraphMatrix) VertexNextNeighbor(rank, currentNeighbor int) int {
	//TODO implement me
	panic("implement me")
}

func (g GraphMatrix) VertexStatus(rank int) int {
	//TODO implement me
	panic("implement me")
}

func (g GraphMatrix) VertexDTime(rank int) int {
	//TODO implement me
	panic("implement me")
}

func (g GraphMatrix) VertexFTime(rank int) int {
	//TODO implement me
	panic("implement me")
}

func (g GraphMatrix) VertexParent(rank int) int {
	//TODO implement me
	panic("implement me")
}

func (g GraphMatrix) VertexPriority(rank int) int {
	//TODO implement me
	panic("implement me")
}

func (g GraphMatrix) Reset() {
	//TODO implement me
	panic("implement me")
}
