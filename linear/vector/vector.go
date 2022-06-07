package vector

import (
	"fmt"

	"gox/utils"
	"math/rand"
	"strconv"
	"time"
)

//向量默认容量
const (
	DefaultCapacity = 2
)

type Vector[T any] struct {
	elems    []T
	size     int
	capacity int
}

func (this *Vector[T]) String() string {
	s := "size:" + strconv.Itoa(this.size) + "\n" +
		"capacity:" + strconv.Itoa(this.capacity) + "\n" +
		"elems:" + fmt.Sprintf("%v", this.elems[:this.Size()])
	return s
}

func NewVector[T any](size int) *Vector[T] {

	initCapacity := DefaultCapacity
	if size > 0 {
		initCapacity = size << 1
	}
	vector := &Vector[T]{
		elems:    make([]T, initCapacity),
		size:     size,
		capacity: initCapacity,
	}
	return vector
}

func (this *Vector[T]) Size() int {
	return this.size
}

func (this *Vector[T]) CopyFromSlice(src []T) {
	lenSrc := len(src)
	this.elems = make([]T, lenSrc<<1)
	this.capacity = lenSrc << 1
	for i := 0; i < lenSrc; i++ {
		this.elems[i] = src[i]
	}
	this.size = lenSrc
}

func (this *Vector[T]) CopyFromVectorSection(src *Vector[T], low, high int) {
	this.CopyFromSlice(src.elems[low:high])
}

func (this *Vector[T]) CopyFromVector(src *Vector[T]) {
	this.CopyFromVectorSection(src, 0, src.Size())
}

func (this *Vector[T]) Clone() *Vector[T] {
	vector := &Vector[T]{}
	vector.CopyFromVector(this)
	return vector
}

func (this *Vector[T]) Get(rank int) T {
	return this.elems[rank]
}

func (this *Vector[T]) Put(rank int, value T) {
	this.elems[rank] = value
}

func (this *Vector[T]) Expand() {
	if this.size < this.capacity {
		return
	}
	oldElems := this.elems
	this.elems = make([]T, this.size<<1)
	this.capacity = this.size << 1
	for i := 0; i < this.size; i++ {
		this.elems[i] = oldElems[i]
	}
}

func (this *Vector[T]) Shrink() {
	//less 1/4 cap,shrink,else not shrink
	if this.capacity<<2 < this.size {
		return
	}
	newElems := make([]T, this.capacity>>1)
	for i := 0; i < this.size; i++ {
		newElems[i] = this.elems[i]
	}
	this.capacity >>= 1
	this.elems = newElems
}
func (this *Vector[T]) Insert(rank int, value T) T {
	this.Expand()
	for i := this.size; i > rank; i-- {
		this.elems[i] = this.elems[i-1]
	}
	this.elems[rank] = value
	this.size++
	return value
}

func (this *Vector[T]) RemoveSection(low, high int) int {
	for high < this.Size() {
		this.elems[low] = this.elems[high]
		high++
		low++
	}
	this.size = low
	return high - low
}

func (this *Vector[T]) Remove(rank int) T {
	removedValue := this.Get(rank)
	this.RemoveSection(rank, rank+1)
	return removedValue
}

func (this *Vector[T]) Append(value T) {
	this.Insert(this.Size(), value)
}

func (this *Vector[T]) Appends(values ...T) {
	for _, v := range values {
		this.Append(v)
	}
}

func (this *Vector[T]) UnSort() {
	this.UnSortSection(0, this.Size())
}

func (this *Vector[T]) UnSortSection(low, high int) {
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := high - 1; i > low; i-- {
		utils.Swap[T](&this.elems[i], &this.elems[low+rand.Int()%(high-low)])
	}
}

func (this *Vector[T]) Traverse(visit func(T) bool) bool {
	for _, v := range this.elems[:this.Size()] {
		if !visit(v) {
			return false
		}
	}
	return true
}

func (this *Vector[T]) Disordered(less func(T, T) bool) int {
	counter := 0
	for i := 1; i < this.Size(); i++ {
		if less(this.Get(i), this.Get(i-1)) {
			counter++
		}
	}
	return counter
}

func (this *Vector[T]) Find(target T, equal func(T, T) bool) int {
	return this.FindSection(target, equal, 0, this.Size())
}

func (this *Vector[T]) FindSection(target T, equal func(T, T) bool, low, high int) int {
	for i := high - 1; i >= low; i-- {
		if equal(this.Get(i), target) {
			return i
		}
	}
	return low - 1
}

//func (this *Vector) Search(target interface{}, less func(interface{}, interface{}) bool) int {
//	return search.BinarySearch(target, this.elems, less)
//}
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
	return oldSize - this.size
}

func (this *Vector[T]) Uniquify(equal func(T, T) bool) int {
	i, j := 0, 1
	for j < this.Size() {
		if !equal(this.Get(i), this.Get(j)) {
			i++
			this.elems[i] = this.elems[j]
		}
		j++
	}
	this.size = this.size - (j - i - 1)
	return j - i - 1
}

//TODO  sort

//TODO  search
