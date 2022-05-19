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
