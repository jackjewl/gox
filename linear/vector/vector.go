//vector包，向量
package vector

import (
	"fmt"
	GoxSearch "gox/search"
	GoxSort "gox/sort"
	"gox/utils"
	"math/rand"
	"time"
)

//constuctor function
func NewVector[T any](size int) *Vector[T] {
	if size < 0 {
		panic("size can negative")
	}
	vector := &Vector[T]{
		elements: make([]T, size),
	}
	return vector
}

//vector
type Vector[T any] struct {
	elements []T //数据区
}

//impl Stringer interface
func (this *Vector[T]) String() string {
	return fmt.Sprintf("%v", this.elements)
}

//size
func (this *Vector[T]) Size() int {
	return len(this.elements)
}

//copy from slice
func (this *Vector[T]) CopyFromSlice(src []T) {
	this.elements = make([]T, len(src))
	copy(this.elements, src)
}

//copy from a section of vector elems ,from low to high
func (this *Vector[T]) CopyFromVectorSection(src *Vector[T], start, end int) {
	this.CopyFromSlice(src.elements[start:end])
}

//copy from all vector elems
func (this *Vector[T]) CopyFromVector(src *Vector[T]) {
	this.CopyFromVectorSection(src, 0, src.Size())
}

//clone a vector same as this vector itself
func (this *Vector[T]) Clone() *Vector[T] {
	vector := NewVector[T](0)
	vector.CopyFromVector(this)
	return vector
}

// get elems by rank
func (this *Vector[T]) Get(rank int) T {
	return this.elements[rank]
}

//put rank value
func (this *Vector[T]) Assign(rank int, value T) {
	this.elements[rank] = value
}

//insert elems value of params  value into rank params rank position
func (this *Vector[T]) Insert(rank int, value T) {

	this.elements = append(this.elements, value)
	copy(this.elements[rank+1:], this.elements[rank:])
	this.elements[rank] = value
}

//remove elems where rank between low and high
func (this *Vector[T]) RemoveSection(low, high int) int {

	copy(this.elements[low:], this.elements[high:])
	this.elements = this.elements[:len(this.elements)-high+low]
	return high - low
}

//remove one elem
func (this *Vector[T]) Remove(rank int) T {
	removedValue := this.Get(rank)
	this.RemoveSection(rank, rank+1)
	return removedValue
}

//append one elem
func (this *Vector[T]) Append(value T) {
	this.elements = append(this.elements, value)
}

//append elems
func (this *Vector[T]) Appends(values ...T) {
	this.elements = append(this.elements, values...)
}

//unsort
func (this *Vector[T]) UnSort() {
	this.unSortSection(0, this.Size())
}

//unsort a section between low and high
func (this *Vector[T]) unSortSection(start, end int) {
	rand.Seed(time.Now().UnixNano())

	for i := end - 1; i > start; i-- {
		utils.Swap[T](&this.elements[i], &this.elements[start+rand.Intn(i-start)])
	}
}

//traverse and  visit each elems
func (this *Vector[T]) Traverse(visit func(T) bool) bool {
	for _, v := range this.elements[:this.Size()] {
		if !visit(v) {
			return false
		}
	}
	return true
}

//disordered
func (this *Vector[T]) Disordered(compare func(T, T) int) bool {
	return this.adjacentInversePairsCount(compare) == 0
}

func (this *Vector[T]) adjacentInversePairsCount(compare func(T, T) int) int {
	counter := 0
	for i := 1; i < this.Size(); i++ {
		if compare(this.Get(i), this.Get(i-1)) < 0 {
			counter++
		}
	}
	return counter
}

//unsort find
func (this *Vector[T]) Find(target T, compare func(T, T) int) int {
	return this.findSection(target, compare, 0, this.Size())
}

//unsort vector find in section
func (this *Vector[T]) findSection(target T, compare func(T, T) int, start, end int) int {
	for i := end - 1; i >= start; i-- {
		if compare(this.Get(i), target) == 0 {
			return i
		}
	}
	return start - 1
}

//unsort deduplicate
func (this *Vector[T]) Deduplicate(compare func(T, T) int) int {
	oldSize := this.Size()
	i := 1
	for i < this.Size() {
		findBeforeResult := this.findSection(this.Get(i), compare, 0, i)
		if findBeforeResult >= 0 {
			this.Remove(findBeforeResult)
		} else {
			i++
		}
	}
	return oldSize - this.Size()
}

//sorted vector uniquify
func (this *Vector[T]) Uniquify(equal func(T, T) bool) int {
	i, j := 0, 1
	for j < this.Size() {
		if !equal(this.Get(i), this.Get(j)) {
			i++
			this.elements[i] = this.elements[j]
		}
		j++
	}
	this.elements = this.elements[:this.Size()-j+i-1]
	return j - i - 1
}

//sort
func (this *Vector[T]) Sort(compare func(T, T) int) {
	GoxSort.QuickSort[T](this.elements, 0, this.Size(), compare)
}

//search
func (this *Vector[T]) Search(target T, compare func(T, T) int) int {
	return GoxSearch.BinarySearch(target, this.elements, compare)
}
