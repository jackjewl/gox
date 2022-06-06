package slice

func BinarySearch(target interface{}, elems []interface{}, less func(interface{}, interface{}) bool) int {
	return BinarySearchSection(target, elems, less, 0, len(elems))
}

func BinarySearchSection(target interface{}, elems []interface{}, less func(interface{}, interface{}) bool, low, high int) int {
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
