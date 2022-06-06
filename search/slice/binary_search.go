package slice

func BinarySearch[T any](target T, elems []T, less func(T, T) bool) int {
	return BinarySearchSection[T](target, elems, less, 0, len(elems))
}

func BinarySearchSection[T any](target T, elems []T, less func(T, T) bool, low, high int) int {
	var mid int
	for low <= high {
		mid = (low + high) >> 1
		if less(target, elems[mid]) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low - 1
}
