package others

import (
	"log"
	"testing"
)

func TestSum(t *testing.T) {
	elems := []int{1, 2, 3, 4, 5}
	log.Println(Sum(elems))
}

func TestSum1(t *testing.T) {
	log.Println(Sum1([]int{1, 2, 3, 4, 5}))
}

func TestSum3(t *testing.T) {
	log.Println(Sum3([]int{1, 2, 3, 4, 5}))
}
