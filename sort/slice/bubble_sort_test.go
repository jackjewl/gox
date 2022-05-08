package slice

import (
	"log"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	elems := []interface{}{23, 5, 40, 3, 2, 1}
	less := func(a interface{}, b interface{}) bool {
		return a.(int) < b.(int)
	}

	BubbleSort(elems, less)
	log.Println(elems...)
}
