package others

import (
	"log"
	"testing"
)

func TestReverse(t *testing.T) {
	elems := []int{1, 2, 3, 4, 5}
	Reverse[int](elems)
	log.Println(elems)
}

func TestReverse1(t *testing.T) {
	elems := []int{1, 2, 3, 4, 5}
	Reverse1[int](elems)
	log.Println(elems)
}

func TestReverse2(t *testing.T) {
	elems := []int{1, 2, 3, 4, 5}
	Reverse2[int](elems)
	log.Println(elems)
}
