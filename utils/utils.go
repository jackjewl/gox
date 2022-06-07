package utils

import (
	"golang.org/x/exp/constraints"
)

func Swap[T any](a *T, b *T) {
	*a, *b = *b, *a
}

func Max[T constraints.Integer | constraints.Float](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

func Sum[T constraints.Integer | constraints.Float](elems []T) T {
	var sum T
	for _, v := range elems {
		sum += v
	}
	return sum
}

//求第n个斐波那契数
func FibNumber(n int) int {
	if n < 2 {
		return n
	}
	pre, next := 0, 1
	for i := 1; i < n; i++ {
		next = next + pre
		pre = next - pre
	}
	return next
}

//切片元素整体左移distance位
func LeftShift[T any](elems []T, distance int) {
	Reverse[T](elems[:distance])
	Reverse[T](elems[distance:])
	Reverse[T](elems)

}

//切片元素整体右移distance位
func RightShift[T any](elems []T, distance int) {
	Reverse[T](elems[len(elems)-distance:])
	Reverse[T](elems[:len(elems)-distance])
	Reverse[T](elems)
}

func Reverse[T any](elems []T) {
	low, high := 0, len(elems)-1
	for low < high {
		Swap[T](&elems[low], &elems[high])
		low++
		high--
	}
}
