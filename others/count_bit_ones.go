package others

func CountBitOnes(elem int) int {
	counter := 0

	for elem > 0 {
		counter += elem & 1
		elem >>= 1
	}
	return counter
}
