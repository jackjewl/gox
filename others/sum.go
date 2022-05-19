package others

func Sum(elems []int) int {
	sum := 0
	for i := 0; i < len(elems); i++ {
		sum += elems[i]
	}
	return sum
}

func Sum1(elems []int) int {
	if len(elems) < 1 {
		return 0
	} else {
		return Sum1(elems[:len(elems)-1]) + elems[len(elems)-1]
	}
}

func Sum3(elems []int) int {
	if len(elems) == 1 {
		return elems[0]
	} else {
		mid := len(elems) >> 1
		return Sum3(elems[:mid]) + Sum3(elems[mid:len(elems)])
	}
}
