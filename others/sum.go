package others

//迭代求数组和
func Sum1(elems []int) int {
	sum := 0
	for i := 0; i < len(elems); i++ {
		sum += elems[i]
	}
	return sum
}

//数组求和递归法，减而治之
func Sum2(elems []int) int {
	if len(elems) < 1 {
		return 0
	} else {
		return Sum1(elems[:len(elems)-1]) + elems[len(elems)-1]
	}
}

//数组求和，分治法
func Sum3(elems []int) int {
	if len(elems) < 1 {
		return 0
	} else if len(elems) == 1 {
		return elems[len(elems)-1]
	} else {
		mid := len(elems) >> 1
		return Sum3(elems[:mid]) + Sum3(elems[mid:len(elems)])
	}
}
