package matrix

type Vertex[T any] struct {
	data      T
	inDegree  int
	outDegree int
	status    int
	dTime     int
	fTime     int
	parent    int
	priority  int
}
