package search

func BinarySearch[T any](target T, elems []T, compare func(T, T) int) int {
	return BinarySearchSection[T](target, elems, compare, 0, len(elems))
}

func BinarySearchSection[T any](target T, elems []T, compare func(T, T) int, start, end int) int {
	var mid int
	for start <= end {
		mid = (start + end) >> 1
		if compare(target, elems[mid]) < 0 {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start - 1
}
