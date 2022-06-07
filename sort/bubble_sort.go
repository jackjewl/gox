package goxSort

func BubbleSort(elems []interface{}, compare func(interface{}, interface{}) bool) {

	high := Bubble(elems, compare)
	for high > 1 {
		high = Bubble(elems[:high], compare)
	}
}

func Bubble(elems []interface{}, less func(interface{}, interface{}) bool) int {
	lastReverseIndex := 0
	for i := 1; i < len(elems); i++ {
		if less(elems[i], elems[i-1]) {
			elems[i], elems[i-1] = elems[i-1], elems[i]
			lastReverseIndex = i
		}
	}
	return lastReverseIndex
}
