package others

import (
	"log"
	"math/rand"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var elems []int
	for i := 0; i < 10; i++ {
		elems = append(elems, rand.Int()%1000)
	}
	less := func(a int, b int) bool {
		return a < b
	}
	BubbleSort[int](elems, less)
	log.Println(elems)
}
