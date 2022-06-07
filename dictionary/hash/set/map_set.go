package set

func NewMapSet[T HashCoder]() *MapSet[T] {
	mapSet := &MapSet[T]{
		mapper: make(map[string]T),
	}
	return mapSet
}

type MapSet[T HashCoder] struct {
	mapper map[string]T
}

func (this *MapSet[T]) Add(element T) bool {
	if _, ok := this.mapper[element.HashCode()]; ok {
		return false
	}
	this.mapper[element.HashCode()] = element
	return true

}
func (this *MapSet[T]) Remove(element T) bool {
	if _, ok := this.mapper[element.HashCode()]; !ok {
		return false
	}
	delete(this.mapper, element.HashCode())
	return true
}
func (this *MapSet[T]) Exist(element T) bool {
	_, ok := this.mapper[element.HashCode()]
	return ok
}
func (this *MapSet[T]) Empty() bool {
	return this.Size() == 0
}
func (this *MapSet[T]) Size() int {
	return len(this.mapper)
}

func (this *MapSet[T]) Union(src MapSet[T]) *MapSet[T] {
	newSet := NewMapSet[T]()
	for _, v := range src.mapper {
		newSet.Add(v)
	}
	for _, v := range this.mapper {
		newSet.Add(v)
	}
	return newSet
}
func (this *MapSet[T]) Intersection(src MapSet[T]) *MapSet[T] {

	newSet := NewMapSet[T]()
	for _, v := range src.mapper {
		if !this.Add(v) {
			newSet.Add(v)
		}
	}

	return newSet
}
func (this *MapSet[T]) Complementary(universe MapSet[T]) *MapSet[T] {
	newSet := NewMapSet[T]()
	for _, v := range this.mapper {
		if universe.Add(v) {
			newSet.Add(v)
		}
	}
	return newSet
}

//this - src
func (this *MapSet[T]) Sub(src MapSet[T]) *MapSet[T] {
	return this.Intersection(src).Complementary(*this)
}

//TODO  并交补的返回值和set接口定义不一样，该怎么统一
