package vector

import (
	search "gox/search/slice"
	sort "gox/sort/slice"
	"gox/utils"
	"math/rand"
)

const (
	DefaultSize = 2
)

type Vector struct {
	elems    []interface{}
	size     int
	capacity int
}

func NewVector(size int) *Vector {
	initSize := DefaultSize
	if size >= 0 {
		initSize = size
	}
	vector := &Vector{
		elems:    make([]interface{}, initSize<<1),
		size:     initSize,
		capacity: initSize << 1,
	}
	return vector
}

func (this *Vector) Size() int {
	return this.size
}

func (this *Vector) CopyFromSlice(src []interface{}) *Vector {
	lenSrc := len(src)
	this.elems = make([]interface{}, lenSrc<<1)
	for i := 0; i < lenSrc; i++ {
		this.elems[i] = src[i]
	}
	this.size = lenSrc
	this.capacity = lenSrc << 1
	return this
}

func (this *Vector) CopyFromVectorSection(src *Vector, low, high int) *Vector {
	return this.CopyFromSlice(src.elems[low:high])
}

func (this *Vector) CopyFromVector(src *Vector) *Vector {
	return this.CopyFromVectorSection(src, 0, this.Size())
}

func (this *Vector) Clone() *Vector {
	vector := &Vector{}
	vector.CopyFromVector(this)
	return vector
}

func (this *Vector) Get(rank int) interface{} {
	return this.elems[rank]
}

func (this *Vector) Put(rank int, value interface{}) *Vector {
	this.elems[rank] = value
	return this
}

func (this *Vector) Expand() *Vector {
	if this.size < this.capacity {
		return this
	}
	oldElems := this.elems
	this.elems = make([]interface{}, this.size<<1)
	this.capacity = this.size << 1
	for i := 0; i < this.size; i++ {
		this.elems[i] = oldElems[i]
	}
	return this
}

func (this *Vector) Shrink() *Vector {
	//less 1/4 cap,shrink,else not shrink
	if this.capacity<<2 < this.size {
		return this
	}
	newElems := make([]interface{}, this.capacity>>1)
	for i := 0; i < this.size; i++ {
		newElems[i] = this.elems[i]
	}
	this.capacity >>= 1
	this.elems = newElems
	return this
}
func (this *Vector) Insert(rank int, value interface{}) interface{} {
	this.Expand()
	for i := this.size; i > rank; i-- {
		this.elems[i] = this.elems[i-1]
	}
	this.elems[rank] = value
	return value
}

func (this *Vector) RemoveSection(low, high int) int {
	for high < this.size {
		this.elems[low] = this.elems[high]
		high++
		low++
		this.size--
	}
	return high - low
}

func (this *Vector) Remove(rank int) interface{} {
	removedValue := this.Get(rank)
	this.RemoveSection(rank, rank+1)
	return removedValue
}

func (this *Vector) Append(value interface{}) *Vector {
	this.Insert(this.Size(), value)
	return this
}

func (this *Vector) Appends(values ...interface{}) *Vector {
	for _, v := range values {
		this.Append(v)
	}
	return this
}

func (this *Vector) Permute() *Vector {
	this.PermuteSection(0, this.Size())
	return this
}

func (this *Vector) PermuteSection(low, high int) *Vector {
	PermuteSlice(this.elems[low:high])
	return this
}
func PermuteSlice(src []interface{}) {
	for i := len(src); i > 1; i-- {
		utils.Swap(&src[rand.Int()%i], &src[i])
	}
}

func (this *Vector) Traverse(visit func(interface{}) bool) bool {
	for _, v := range this.elems {
		if !visit(v) {
			return false
		}
	}
	return true
}

func (this *Vector) Disordered(less func(interface{}, interface{}) bool) bool {
	return this.AdjacentReverseOrderPairCount(less) == 0
}

//get counter of adjacent reverse order pair,example1,5,4,1,return 2,
func (this *Vector) AdjacentReverseOrderPairCount(less func(interface{}, interface{}) bool) int {
	count := 0
	for i := 1; i < this.Size(); i++ {
		if !less(this.elems[i-1], this.elems[i]) {
			count++
		}
	}
	return count
}

func (this *Vector) Find(target interface{}, equal func(interface{}, interface{}) bool) int {
	for i := this.Size(); i > 0; i-- {
		if equal(this.Get(i), target) {
			return i
		}
	}
	return -1
}

func (this *Vector) Search(target interface{}, less func(interface{}, interface{}) bool) int {
	return search.BinarySearch(target, this.elems, less)
}
func (this *Vector) Deduplicate(equal func(interface{}, interface{}) bool) int {
	oldSize := this.Size()
	for i := 1; i < this.Size(); i++ {
		findBeforeResult := this.Find(this.Get(i), equal)
		if findBeforeResult > 0 {
			this.Remove(findBeforeResult)
		}
	}
	return oldSize - this.size
}

func (this *Vector) Uniquify(equal func(interface{}, interface{}) bool) int {
	i, j := 0, 1
	for j < this.Size() {
		if !equal(this.Get(i), this.Get(j)) {
			i++
			this.elems[i] = this.elems[j]
		}
		j++
	}
	return j - i - 1
}

func (this *Vector) Sort(less func(interface{}, interface{}) bool) *Vector {
	sort.MergeSort(this.elems, less)
	return this
}
