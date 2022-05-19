package others

import (
	"gox/utils"
)

func Reverse[T any](elems []T) {
	if len(elems) > 1 {
		utils.Swap[T](&elems[0], &elems[len(elems)-1])
		Reverse[T](elems[1 : len(elems)-1])
	}
}

func Reverse1[T any](elems []T) {
	low, high := 0, len(elems)-1
next:
	if low < high {
		utils.Swap[T](&elems[low], &elems[high])
		low++
		high--
		goto next
	}
}

func Reverse2[T any](elems []T) {
	low, high := 0, len(elems)-1
	for low < high {
		utils.Swap[T](&elems[low], &elems[high])
		low++
		high--
	}
}
