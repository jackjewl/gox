package others

import (
	"gox/utils"
)

//反转数组，递归法
func Reverse1[T any](elems []T) {
	if len(elems) > 1 {
		utils.Swap[T](&elems[0], &elems[len(elems)-1])
		Reverse1[T](elems[1 : len(elems)-1])
	}
}

//反转数组，用goto消除尾递归
func Reverse2[T any](elems []T) {
	low, high := 0, len(elems)-1
next:
	if low < high {
		utils.Swap[T](&elems[low], &elems[high])
		low++
		high--
		goto next
	}
}

//反转数组，用循环消除尾递归
func Reverse3[T any](elems []T) {
	low, high := 0, len(elems)-1
	for low < high {
		utils.Swap[T](&elems[low], &elems[high])
		low++
		high--
	}
}
