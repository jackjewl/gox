package vector

import (
	"log"
	"testing"
)

func TestNewVector(t *testing.T) {
	v := NewVector[int](0)
	v1 := NewVector[int](1)
	v2 := NewVector[int](3)
	log.Println(v)
	log.Println(v1)
	log.Println(v2)
}

func TestVector_Append(t *testing.T) {
	v1 := NewVector[int](0)
	v1.Append(11)
	v1.Append(11)
	v1.Append(11)
	log.Println(v1)

}

func TestVector_Appends(t *testing.T) {
	v1 := NewVector[int](0)
	v1.Appends(1, 2, 3, 4, 5)
	log.Println(v1)

}

func TestVector_Clone(t *testing.T) {
	v1 := NewVector[int](0)
	v1.Appends(1, 2, 3, 4)
	v2 := v1.Clone()
	log.Println(v1)
	log.Println(v2)
}

func TestVector_CopyFromSlice(t *testing.T) {
	v1 := NewVector[int](0)
	v1.CopyFromSlice([]int{1, 2, 3, 4, 5})
	log.Println(v1)

}

func TestVector_CopyFromVector(t *testing.T) {
	v1 := NewVector[int](0)
	v1.Appends(1, 2, 3, 4, 5)
	v2 := NewVector[int](0)
	v2.CopyFromVector(v1)
	log.Println(v1)
	log.Println(v2)
}

func TestVector_CopyFromVectorSection(t *testing.T) {
	v1 := NewVector[int](0)
	v1.Appends(1, 2, 3, 4, 5)
	v2 := NewVector[int](0)
	v2.CopyFromVectorSection(v1, 2, 5)
	log.Println(v1)
	log.Println(v2)

}

func TestVector_Deduplicate(t *testing.T) {
	v1 := NewVector[int](0)
	v1.Appends(1, 1, 1, 2, 2, 2, 3, 4, 5, 6)

	equal := func(a, b int) bool {
		return a == b
	}
	v1.Deduplicate(equal)
	log.Println(v1)

}

func TestVector_Disordered(t *testing.T) {
	v1 := NewVector[int](0)
	v1.Appends(1, 1, 1, 2, 2, 2, 3, 4, 5, 6)

	less := func(a, b int) bool {
		return a < b
	}
	log.Println(v1.Disordered(less) == 0)

	v2 := NewVector[int](0)
	v2.Appends(11, 1, 1, 2, 2, 2, 3, 4, 5, 6)
	log.Println(v2.Disordered(less) == 0)

}

func TestVector_Expand(t *testing.T) {

	v1 := NewVector[int](0)
	log.Println(v1)
	v1.Expand()
	log.Println(v1)

	v1.Appends(1, 2, 3, 4)
	log.Println(v1)
	v1.Expand()
	log.Println(v1)

}

func TestVector_Find(t *testing.T) {
	v1 := NewVector[int](0)
	v1.Appends(1, 2, 3, 4, 5)
	equal := func(a, b int) bool {
		return a == b
	}
	log.Println(v1.Find(2, equal))
	log.Println(v1.Find(9, equal))
	log.Println(v1.Find(5, equal))

}

func TestVector_FindSection(t *testing.T) {
	v1 := NewVector[int](0)
	v1.Appends(1, 2, 3, 4, 5)
	equal := func(a, b int) bool {
		return a == b
	}
	log.Println(v1.FindSection(3, equal, 1, 4))

}

func TestVector_Get(t *testing.T) {
	v1 := NewVector[int](0)
	v1.Appends(1, 2, 3, 4)
	log.Println(v1.Get(2))
}

func TestVector_Insert(t *testing.T) {
	v1 := NewVector[int](0)
	v1.Insert(v1.Size(), 11)
	v1.Insert(v1.Size(), 11)
	v1.Insert(v1.Size(), 11)
	v1.Insert(v1.Size(), 11)
	v1.Insert(v1.Size(), 11)
	log.Println(v1)
}

func TestVector_Put(t *testing.T) {

	v1 := NewVector[int](0)
	v1.Appends(1, 2, 3, 4)
	v1.Put(1, 10)
	log.Println(v1)

}

func TestVector_Remove(t *testing.T) {
	v1 := NewVector[int](0)
	v1.Appends(1, 2, 3, 4, 5, 6, 7, 8, 9)
	v1.Remove(3)
	log.Println(v1)
}

func TestVector_RemoveSection(t *testing.T) {
	v1 := NewVector[int](0)
	v1.Appends(1, 2, 3, 4, 5, 6, 7, 8, 9)
	v1.RemoveSection(2, 5)
	log.Println(v1)
}

func TestVector_Shrink(t *testing.T) {
	v1 := NewVector[int](0)
	v1.Appends(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	log.Println(v1)
	v1.RemoveSection(2, 9)
	log.Println(v1)
	v1.Shrink()
	log.Println(v1)
}

func TestVector_Size(t *testing.T) {
	v1 := NewVector[int](0)
	log.Println(v1.Size())

}

func TestVector_String(t *testing.T) {
	v1 := NewVector[int](0)
	log.Println(v1.String())
}

func TestVector_Traverse(t *testing.T) {

	visit := func(a int) bool {
		log.Println(a)
		return true
	}
	v1 := NewVector[int](0)
	v1.Appends(1, 2, 3, 4, 5)
	v1.Traverse(visit)

}

func TestVector_UnSort(t *testing.T) {

	v1 := NewVector[int](0)
	v1.Appends(1, 2, 3, 4, 5, 6, 7)
	v1.UnSort()
	log.Println(v1)

}

func TestVector_UnSortSection(t *testing.T) {
	v1 := NewVector[int](0)
	v1.Appends(1, 2, 3, 4, 5, 6, 7)
	v1.UnSortSection(2, 5)
	log.Println(v1)
}

func TestVector_Uniquify(t *testing.T) {
	v1 := NewVector[int](0)
	v1.Appends(1, 1, 1, 1, 2, 2, 2, 3, 4, 5, 5, 6, 6, 6)
	equal := func(a, b int) bool {
		return a == b
	}
	v1.Uniquify(equal)
	log.Println(v1)

}

func TestVector_BinarySearchSectionA(t *testing.T) {
	v1 := NewVector[int](0)
	v1.Appends(1, 2, 3, 4, 5, 6, 7, 8, 9)
	compare := func(a, b int) int {
		return a - b
	}
	r := v1.BinarySearchSectionA(9, compare, 0, 9)
	log.Println(r)
	r1 := v1.BinarySearchSectionA(8, compare, 0, 9)
	log.Println(r1)
	r2 := v1.BinarySearchSectionA(7, compare, 0, 9)
	log.Println(r2)
	r3 := v1.BinarySearchSectionA(6, compare, 0, 9)
	log.Println(r3)
	r4 := v1.BinarySearchSectionA(5, compare, 0, 9)
	log.Println(r4)

}

func TestVector_FibonacciSearchSection(t *testing.T) {
	v1 := NewVector[int](0)
	v1.Appends(1, 2, 3, 4, 5, 6, 7, 8, 9)
	compare := func(a, b int) int {
		return a - b
	}
	r := v1.FibonacciSearchSection(9, compare, 0, 9)
	log.Println(r)
	r1 := v1.FibonacciSearchSection(8, compare, 0, 9)
	log.Println(r1)
	r2 := v1.FibonacciSearchSection(7, compare, 0, 9)
	log.Println(r2)
	r3 := v1.FibonacciSearchSection(0, compare, 0, 9)
	log.Println(r3)
	r4 := v1.FibonacciSearchSection(1, compare, 0, 9)
	log.Println(r4)
}

func TestVector_BinarySearch(t *testing.T) {
	v1 := NewVector[int](0)
	v1.Appends(1, 2, 3, 4, 5, 6, 7, 8, 9)
	compare := func(a, b int) int {
		return a - b
	}
	r := v1.BinarySearch(9, compare)
	log.Println(r)
	r1 := v1.BinarySearch(1, compare)
	log.Println(r1)
}

func TestVector_BubbleSort(t *testing.T) {
	v1 := NewVector[int](0)
	v1.Appends(11, 2, 3, 4, 35, 6, 7, 48, 9)
	less := func(a, b int) bool {
		return a < b
	}
	v1.BubbleSort(less)
	log.Println(v1)
}

func TestVector_MergeSort(t *testing.T) {
	v1 := NewVector[int](0)
	v1.Appends(11, 2, 3, 4, 35, 6, 7, 48, 9)
	less := func(a, b int) bool {
		return a < b
	}
	v1.MergeSort(less)
	log.Println(v1)
}
