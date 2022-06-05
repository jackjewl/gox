package slice

import (
	"gox/utils"
)

func ShellSort[T any](elems []T, compare func(T, T) int) {
	for width := len(elems) >> 1; width > 0; width >>= 1 {
		ShellWithWidth[T](elems, compare, width)
	}
}

func ShellWithWidth[T any](elems []T, compare func(T, T) int, width int) {

	for groupStart := 0; groupStart <= width; groupStart++ {
		for outIndex := groupStart + width; outIndex < len(elems); outIndex += width {
			key := elems[outIndex]
			innerIndex := outIndex - width
			for innerIndex >= 0 && compare(elems[innerIndex], key) > 0 {
				utils.Swap[T](&elems[innerIndex+width], &elems[innerIndex])
				innerIndex -= width
			}
			elems[innerIndex+width] = key
		}

	}

}

//[22 2 11 64 2 34 60 3]
