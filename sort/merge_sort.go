package goxSort

func MergeSort(elems []interface{}, less func(interface{}, interface{}) bool) {
	MergeSortPart(elems, less, 0, len(elems))
}

func MergeSortPart(elems []interface{}, less func(interface{}, interface{}) bool, low, high int) {
	if low-high < 2 {
		return
	}
	mid := (low + high) >> 1
	MergeSortPart(elems, less, low, mid)
	MergeSortPart(elems, less, mid, high)
	Merge(elems, less, low, mid, high)
}

func Merge(elems []interface{}, less func(interface{}, interface{}) bool, low, mid, high int) {
	tempSlice := make([]interface{}, high-low)
	leftIndex, rightIndex, tempIndex := low, mid, 0
	for leftIndex < mid && rightIndex < high {
		if less(elems[rightIndex], elems[leftIndex]) {
			tempSlice[tempIndex] = elems[rightIndex]
			rightIndex++
		} else {
			tempSlice[tempIndex] = elems[leftIndex]
			leftIndex++
		}
		tempIndex++
	}
	if leftIndex < mid-1 {
		tempSlice[tempIndex] = elems[leftIndex]
	}
	if rightIndex < high-1 {
		tempSlice[tempIndex] = elems[rightIndex]
	}
	for i := 0; i < len(tempSlice); i++ {
		elems[low+i] = tempSlice[i]
	}
}
