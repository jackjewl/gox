package others

//冰雹数序列的长度
func Hailstone(n int) int {
	length := 1

	for 1 < n {
		if n%2 == 0 {
			n = 3*n + 1
		} else {
			n /= 2
		}
		length++
	}
	return length
}
