package adjacent_list

import "gox/vector"

type GraphAdjacentList[VT any, ET any] struct {
	vertexes vector.Vector[Vertex[VT, ET]]
	edgesNum int
}

func (this *GraphAdjacentList[VT, ET]) EdgesNum() int {
	return this.edgesNum
}

func (this *GraphAdjacentList[VT, ET]) EdgeExist(v, u int) bool {

	target := Edge[ET]{
		edgeHeaderIndex: u,
	}

	equal := func(a Edge[ET], b Edge[ET]) bool {
		return a.edgeHeaderIndex == b.edgeHeaderIndex
	}
	return this.vertexes.Get(v).outEdges.Find(target, equal) != this.vertexes.Get(v).outEdges.First()
}

func (this *GraphAdjacentList[VT, ET]) EdgeInsert(v, u int) {

	vertex := this.vertexes.Get(v)
	vertex.outEdges.InsertAsLast(Edge[ET]{
		edgeHeaderIndex: u,
	})
	this.vertexes.Put(v, vertex)
}

func (this *GraphAdjacentList[VT, ET]) EdgeRemove(v, u int) {
	target := Edge[ET]{
		edgeHeaderIndex: u,
	}

	equal := func(a Edge[ET], b Edge[ET]) bool {
		return a.edgeHeaderIndex == b.edgeHeaderIndex
	}
	removedEdgePtr := this.vertexes.Get(v).outEdges.Find(target, equal)
	vertex := this.vertexes.Get(v)
	vertex.outEdges.Remove(removedEdgePtr)
	this.vertexes.Put(v, vertex)

}

func (this *GraphAdjacentList[VT, ET]) EdgeStatus(v, u int) int {
	target := Edge[ET]{
		edgeHeaderIndex: u,
	}
	equal := func(a Edge[ET], b Edge[ET]) bool {
		return a.edgeHeaderIndex == b.edgeHeaderIndex
	}
	return this.vertexes.Get(v).outEdges.Find(target, equal).Data().status
}

func (this *GraphAdjacentList[VT, ET]) EdgeData(v, u int) ET {
	target := Edge[ET]{
		edgeHeaderIndex: u,
	}
	equal := func(a Edge[ET], b Edge[ET]) bool {
		return a.edgeHeaderIndex == b.edgeHeaderIndex
	}
	return this.vertexes.Get(v).outEdges.Find(target, equal).Data().data

}

func (this *GraphAdjacentList[VT, ET]) Weight(v, u int) int {
	target := Edge[ET]{
		edgeHeaderIndex: u,
	}
	equal := func(a Edge[ET], b Edge[ET]) bool {
		return a.edgeHeaderIndex == b.edgeHeaderIndex
	}
	return this.vertexes.Get(v).outEdges.Find(target, equal).Data().weight

}

func (this *GraphAdjacentList[VT, ET]) VertexNum() int {
	return this.vertexes.Size()
}

func (this *GraphAdjacentList[VT, ET]) VertexInsert(vertexData VT) {

	this.vertexes.Append(Vertex[VT, ET]{
		data: vertexData,
	})
}

func (this *GraphAdjacentList[VT, ET]) VertexRemove(rank int) {

	for i := 0; i < this.vertexes.Size(); i++ {
		this.EdgeRemove(rank, i)
	}
	this.vertexes.Remove(rank)
}

func (this *GraphAdjacentList[VT, ET]) VertexInDegree(rank int) int {

	return this.vertexes.Get(rank).inEdges.Size()
}

func (this *GraphAdjacentList[VT, ET]) VertexOutDegree(rank int) int {
	return this.vertexes.Get(rank).outEdges.Size()
}

func (this *GraphAdjacentList[VT, ET]) VertexFirstNeighbor(rank int) int {
	return this.vertexes.Get(rank).outEdges.First().Data().edgeHeaderIndex
}

func (this *GraphAdjacentList[VT, ET]) VertexNextNeighbor(rank, currentNeighbor int) int {
	target := Edge[ET]{
		edgeHeaderIndex: currentNeighbor,
	}

	equal := func(a Edge[ET], b Edge[ET]) bool {
		return a.edgeHeaderIndex == b.edgeHeaderIndex
	}
	curr := this.vertexes.Get(rank).outEdges.Find(target, equal)
	return curr.Succ().Data().edgeHeaderIndex
}

func (this *GraphAdjacentList[VT, ET]) VertexStatus(rank int) int {
	return this.vertexes.Get(rank).status
}

func (this *GraphAdjacentList[VT, ET]) VertexDTime(rank int) int {
	return this.vertexes.Get(rank).dTime
}

func (this *GraphAdjacentList[VT, ET]) VertexFTime(rank int) int {
	return this.vertexes.Get(rank).fTime
}

func (this *GraphAdjacentList[VT, ET]) VertexParent(rank int) int {
	return this.vertexes.Get(rank).parent
}

func (this *GraphAdjacentList[VT, ET]) VertexPriority(rank int) int {
	return this.vertexes.Get(rank).priority
}

//所有边和顶点的状态复位
func (this *GraphAdjacentList[VT, ET]) Reset() {

	for i := 0; i < this.vertexes.Size(); i++ {
		vertexe := this.vertexes.Get(i)
		vertexe.status = 1
		for ptr := vertexe.inEdges.First(); ptr != vertexe.inEdges.Trailer(); ptr = ptr.Succ() {
			//TODO
			//将每条边的状态修改
		}
		this.vertexes.Put(i, vertexe)
	}

}
