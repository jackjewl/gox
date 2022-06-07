//vector包，向量
package vector

import (
	"fmt"
<<<<<<< HEAD:vector/vector.go
	searchSlice "gox/search/slice"
	"gox/sort/slice"
=======

>>>>>>> refact:linear/vector/vector.go
	"gox/utils"
	"math/rand"
	"time"
)

//constuctor function
func NewVector[T any](size int) *Vector[T] {
	if size < 0 {
		panic("not negative")
	}
	vector := &Vector[T]{
		elems: make([]T, size),
	}
	return vector
}

//vector
type Vector[T any] struct {
	elems []T //数据区
}

//impl Stringer interface
func (this *Vector[T]) String() string {
	s := "elems:" + fmt.Sprintf("%v", this.elems[:this.Size()])
	return s
}

//size
func (this *Vector[T]) Size() int {
	return len(this.elems)
}

//copy from slice
func (this *Vector[T]) CopyFromSlice(src []T) {
	this.elems = make([]T, len(src))
	copy(this.elems, src)
}

//copy from a section of vector elems ,from low to high
func (this *Vector[T]) CopyFromVectorSection(src *Vector[T], low, high int) {
	this.CopyFromSlice(src.elems[low:high])
}

//copy from all vector elems
func (this *Vector[T]) CopyFromVector(src *Vector[T]) {
	this.CopyFromVectorSection(src, 0, src.Size())
}

//clone a vector same as this vector itself
func (this *Vector[T]) Clone() *Vector[T] {
	vector := &Vector[T]{}
	vector.CopyFromVector(this)
	return vector
}

// get elems by rank
func (this *Vector[T]) Get(rank int) T {
	return this.elems[rank]
}

//put rank value
func (this *Vector[T]) Put(rank int, value T) {
	this.elems[rank] = value
}

//insert elems value of params  value into rank params rank position
func (this *Vector[T]) Insert(rank int, value T) {

	this.elems = append(this.elems, value)
	copy(this.elems[rank+1:], this.elems[rank:])
	this.elems[rank] = value
}

//remove elems where rank between low and high
func (this *Vector[T]) RemoveSection(low, high int) int {

	copy(this.elems[low:], this.elems[high:])
	this.elems = this.elems[:len(this.elems)-high+low]
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
	this.elems = append(this.elems, value)
}

//append elems
func (this *Vector[T]) Appends(values ...T) {
	this.elems = append(this.elems, values...)
}

//unsort
func (this *Vector[T]) UnSort() {
	this.UnSortSection(0, this.Size())
}

//unsort a section between low and high
func (this *Vector[T]) UnSortSection(low, high int) {
	rand.Seed(time.Now().UnixNano())

	for i := high - 1; i > low; i-- {
		utils.Swap[T](&this.elems[i], &this.elems[low+rand.Intn(i-low)])
	}
}

//traverse and  visit each elems
func (this *Vector[T]) Traverse(visit func(T) bool) bool {
	for _, v := range this.elems[:this.Size()] {
		if !visit(v) {
			return false
		}
	}
	return true
}

//disordered
func (this *Vector[T]) Disordered(less func(T, T) bool) bool {
	counter := 0
	for i := 1; i < this.Size(); i++ {
		if less(this.Get(i), this.Get(i-1)) {
			counter++
		}
	}
	return counter == 0
}

//unsort find
func (this *Vector[T]) Find(target T, equal func(T, T) bool) int {
	return this.FindSection(target, equal, 0, this.Size())
}

//unsort vector find in section
func (this *Vector[T]) FindSection(target T, equal func(T, T) bool, low, high int) int {
	for i := high - 1; i >= low; i-- {
		if equal(this.Get(i), target) {
			return i
		}
	}
	return low - 1
}

//unsort deduplicate
func (this *Vector[T]) Deduplicate(equal func(T, T) bool) int {
	oldSize := this.Size()
	i := 1
	for i < this.Size() {
		findBeforeResult := this.FindSection(this.Get(i), equal, 0, i)
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
			this.elems[i] = this.elems[j]
		}
		j++
	}
	this.elems = this.elems[:this.Size()-j+i-1]
	return j - i - 1
}

<<<<<<< HEAD:vector/vector.go
//sort
func (this *Vector[T]) Sort(compare func(T, T) int) {
	slice.QuickSort[T](this.elems, 0, this.Size(), compare)
}

//search

func (this *Vector[T]) Search(target T, less func(T, T) bool) int {

	return searchSlice.BinarySearch(target, this.elems, less)
}
=======
//TODO  sort

//TODO  search
>>>>>>> refact:linear/vector/vector.go
