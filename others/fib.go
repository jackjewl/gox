package others

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-2) + Fib(n-1)
}

func Fib1(n int) int {
	if n < 2 {
		return n
	}

	pre, next := 0, 1

	for i := 1; i < n; i++ {
		next = next + pre
		pre = next - pre
	}
	return next

}
