package list

import (
	"log"
	"testing"
)

func TestList_Clear(t *testing.T) {
	list := NewList[int]()
	list.InsertAsLast(1)
	list.InsertAsLast(2)
	list.InsertAsLast(3)
	log.Println(list)
	list.Clear()
	log.Println(list)
}

func TestList_CopyList(t *testing.T) {
	list := NewList[int]()
	list.InsertAsLast(1)
	list.InsertAsLast(2)
	list.InsertAsLast(3)
	log.Println(list)
	list1 := NewList[int]()
	list1.CopyList(list)
	log.Println(list1)
}

func TestList_CopyNodes(t *testing.T) {
	list := NewList[int]()
	list.InsertAsLast(1)
	list.InsertAsLast(2)
	list.InsertAsLast(3)
	log.Println(list)
	list1 := NewList[int]()
	log.Println(list1)
	list1.CopyNodes(list.First(), 3)
	log.Println(list1)
}

func TestList_CopyNodesSection(t *testing.T) {

	list := NewList[int]()
	list.InsertAsLast(1)
	list.InsertAsLast(2)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	log.Println(list)
	list1 := NewList[int]()
	log.Println(list1)
	list1.CopyNodesSection(list, 2, 5)
	log.Println(list1)

}

func TestList_Deduplicated(t *testing.T) {
	list := NewList[int]()
	list.InsertAsLast(1)
	list.InsertAsLast(2)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	log.Println(list)
	equal := func(a, b int) bool {
		return a == b
	}
	list.Deduplicated(equal)
	log.Println(list)
}

func TestList_Disordered(t *testing.T) {
	list := NewList[int]()
	list.InsertAsLast(10)
	list.InsertAsLast(2)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	log.Println(list)
	less := func(a, b int) bool {
		return a < b
	}
	r := list.Disordered(less)
	log.Println(r)
}

func TestList_Empty(t *testing.T) {
	list := NewList[int]()
	log.Println(list.Empty())
	list.InsertAsLast(10)
	list.InsertAsLast(2)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	log.Println(list.Empty())
}

func TestList_Find(t *testing.T) {
	list := NewList[int]()
	list.InsertAsLast(10)
	list.InsertAsLast(2)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	equal := func(a, b int) bool {
		return a == b
	}
	log.Println(list.Find(10, equal).data)
	log.Println(list.Find(3, equal).data)

}

func TestList_FindSection(t *testing.T) {
	list := NewList[int]()
	list.InsertAsLast(10)
	list.InsertAsLast(2)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	equal := func(a, b int) bool {
		return a == b
	}
	log.Println(list.FindSection(10, 7, list.Last(), equal).data)
}

func TestList_First(t *testing.T) {
	list := NewList[int]()
	log.Println(list.First().data)
	log.Println(list.Last().data)
	list.InsertAsLast(10)
	list.InsertAsLast(2)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(33)
	log.Println(list.First().data)
	log.Println(list.Last().data)

}

func TestList_Get(t *testing.T) {
	list := NewList[int]()
	list.InsertAsLast(10)
	list.InsertAsLast(2)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(33)
	log.Println(list.Get(0).data, list.Get(1).data, list.Get(2).data)

}

func TestList_Header(t *testing.T) {
	list := NewList[int]()
	list.InsertAsLast(10)
	list.InsertAsLast(2)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(33)
	log.Println(list.Header().data)
	//log.Println(list.Get(0).data, list.Get(1).data, list.Get(2).data)

}

func TestList_InsertAfter(t *testing.T) {
	list := NewList[int]()
	list.InsertAsLast(10)
	list.InsertAsLast(2)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(33)
	list.InsertAfter(list.First(), 222)
	log.Println(list)
}

func TestList_InsertAsFirst(t *testing.T) {
	list := NewList[int]()
	list.InsertAsLast(10)
	list.InsertAsLast(2)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(33)
	log.Println(list)
	list.InsertAsFirst(111)
	log.Println(list)
}

func TestList_InsertAsLast(t *testing.T) {
	list := NewList[int]()
	list.InsertAsLast(10)
	list.InsertAsLast(2)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(33)
	log.Println(list)
}

func TestList_InsertBefroe(t *testing.T) {
	list := NewList[int]()
	list.InsertAsLast(10)
	list.InsertAsLast(2)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(33)
	list.InsertBefore(list.Get(2), 444)
	log.Println(list)
}

func TestList_InsertionSort(t *testing.T) {
	list := NewList[int]()
	list.InsertAsLast(10)
	list.InsertAsLast(2)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(33)
	log.Println(list)
	compare := func(a, b int) int {
		return a - b
	}
	list.InsertionSort(list.First(), list.Size(), compare)
	log.Println(list)
}

func TestList_Last(t *testing.T) {
	list := NewList[int]()
	list.InsertAsLast(10)
	list.InsertAsLast(2)
	list.InsertAsLast(3)
	list.InsertAsLast(32)
	log.Println(list.Last().data)

}

//TODO error
func TestList_Merge(t *testing.T) {
	compare := func(a int, b int) int {
		return a - b
	}
	list := NewList[int]()
	list.InsertAsLast(10)
	list.InsertAsLast(11)
	list.InsertAsLast(12)
	list.InsertAsLast(4)
	list.InsertAsLast(5)
	list.InsertAsLast(6)
	list.InsertAsLast(11)
	list.Merge(list.First(), list.Get(3), list.Last(), compare)
	log.Println(list)

}

//TODO error
func TestList_MergeSort(t *testing.T) {

	list := NewList[int]()
	list.InsertAsLast(100)
	list.InsertAsLast(2)
	list.InsertAsLast(32)
	log.Println(list)
	compare := func(a int, b int) int {
		return a - b
	}
	list.MergeSort(compare)
	log.Println(list)

}

//TODO error
func TestList_MergeSortR(t *testing.T) {

}

func TestList_Remove(t *testing.T) {
	list := NewList[int]()
	list.InsertAsLast(10)
	list.InsertAsLast(2)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(33)
	log.Println(list)
	list.Remove(list.Get(1))
	log.Println(list)
}

func TestList_Search(t *testing.T) {
	list := NewList[int]()
	list.InsertAsLast(10)
	list.InsertAsLast(2)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(33)
	compare := func(a, b int) int {
		return a - b
	}
	log.Println(list.Search(3, list.Size(), list.Last(), compare).data)
	log.Println(list.Search(3, list.Size(), list.Last(), compare).succ.data)
}

func TestList_SelectMax(t *testing.T) {
	list := NewList[int]()
	list.InsertAsLast(10)
	list.InsertAsLast(2)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(33)
	compare := func(a, b int) int {
		return a - b
	}
	println(list.SelectMax(list.First(), list.size, compare).data)
}

//TODO error
func TestList_SelectionSort(t *testing.T) {
	list := NewList[int]()
	list.InsertAsLast(10)
	list.InsertAsLast(2)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(33)
	log.Println(list)
	compare := func(a, b int) int {
		return a - b
	}
	list.SelectionSort(list.First(), list.Size(), compare)
	log.Println(list)
}

func TestList_Size(t *testing.T) {
	list := NewList[int]()
	list.InsertAsLast(10)
	list.InsertAsLast(2)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(33)
	log.Println(list.Size())
}

func TestList_Trailer(t *testing.T) {
	list := NewList[int]()
	list.InsertAsLast(10)
	list.InsertAsLast(2)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(33)
	log.Println(list.Trailer().Pred().Data())

}

func TestList_Traverse(t *testing.T) {
	visit := func(a int) bool {
		log.Println("visit :", a)
		return true
	}
	list := NewList[int]()
	list.InsertAsLast(10)
	list.InsertAsLast(2)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(33)
	list.Traverse(visit)

}

func TestList_Uniquify(t *testing.T) {
	list := NewList[int]()
	list.InsertAsLast(1)
	list.InsertAsLast(2)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(3)
	list.InsertAsLast(33)
	list.InsertAsLast(33)
	list.InsertAsLast(33)
	list.InsertAsLast(33)
	list.InsertAsLast(53)
	log.Println(list)
	equal := func(a, b int) bool {
		return a == b
	}
	list.Uniquify(equal)
	log.Println(list)
}

func TestNewList(t *testing.T) {
	list := NewList[int]()
	log.Println(list, list.Header(), list.Trailer(), list.First(), list.Last(), list.Size())
}
