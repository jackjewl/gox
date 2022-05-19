package others

import (
	"log"
	"testing"
)

func TestRightShift(t *testing.T) {
	elems := []int{1, 2, 3, 4, 5}
	RightShift[int](elems, 2)
	log.Println(elems)
}

func TestLeftShift(t *testing.T) {
	elems := []int{1, 2, 3, 4, 5}
	LeftShift[int](elems, 2)
	log.Println(elems)
}
