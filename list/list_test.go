package list

import (
	"log"
	"testing"
)

func TestList_Clear(t *testing.T) {

}

func TestList_CopyList(t *testing.T) {

}

func TestList_CopyNodes(t *testing.T) {

}

func TestList_CopyNodesSection(t *testing.T) {

}

func TestList_Deduplicated(t *testing.T) {

}

func TestList_Disordered(t *testing.T) {

}

func TestList_Empty(t *testing.T) {

}

func TestList_Find(t *testing.T) {

}

func TestList_FindSection(t *testing.T) {

}

func TestList_First(t *testing.T) {

}

func TestList_Get(t *testing.T) {

}

func TestList_Header(t *testing.T) {

}

func TestList_InsertAfter(t *testing.T) {

}

func TestList_InsertAsFirst(t *testing.T) {

}

func TestList_InsertAsLast(t *testing.T) {

}

func TestList_InsertBefroe(t *testing.T) {

}

func TestList_InsertionSort(t *testing.T) {

}

func TestList_Last(t *testing.T) {

}

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

func TestList_MergeSort(t *testing.T) {

	list := NewList[int]()
	list.InsertAsLast(100)
	list.InsertAsLast(2)
	list.InsertAsLast(32)
	list.InsertAsLast(400)
	list.InsertAsLast(55)
	list.InsertAsLast(55)
	list.InsertAsLast(155)
	list.InsertAsLast(255)
	log.Println(list)
	compare := func(a int, b int) int {
		return a - b
	}
	list.MergeSort(compare)
	log.Println(list)

}

func TestList_MergeSortR(t *testing.T) {

}

func TestList_Remove(t *testing.T) {

}

func TestList_Search(t *testing.T) {

}

func TestList_SelectMax(t *testing.T) {

}

func TestList_SelectionSort(t *testing.T) {

}

func TestList_Size(t *testing.T) {

}

func TestList_Trailer(t *testing.T) {

}

func TestList_Traverse(t *testing.T) {

}

func TestList_Uniquify(t *testing.T) {

}

func TestNewList(t *testing.T) {

}
