package others

func Hailstone(n int) []int {
	if n <= 1 {
		return []int{1}
	}

	if n%2 == 0 {
		a := []int{n}
		a = append(a, Hailstone(n/2)...)
		return a
	} else {
		a := []int{n}
		a = append(a, Hailstone(3*n+1)...)
		return a
	}
}

func HailstoneLength(n int) int {
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
