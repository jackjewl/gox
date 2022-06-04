package utils

import (
	"log"
	"testing"
)

func TestSwap(t *testing.T) {
	a := make([]int, 3, 3)
	a[0] = 1
	a[1] = 2
	a[2] = 3
	SS(a)
	log.Println(a)

}

func SS(a []int) {
	a[1] = 111
	a = append(a, 1, 2, 3)
}
