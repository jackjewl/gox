package others

import (
	"gox/utils"
)

func BubbleSort[T any](elems []T, Less func(T, T) bool) {
	end := len(elems)
	sorted := false
	for !sorted {
		sorted = true
		for i := 1; i < end; i++ {
			if Less(elems[i], elems[i-1]) {
				utils.Swap(&elems[i], &elems[i-1])
				sorted = false
			}
		}
		end--
	}
}
