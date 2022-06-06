package others

//递归版本，低效
func Fib1(n int) int {
	if n < 2 {
		return n
	}
	return Fib1(n-2) + Fib1(n-1)
}

//动态规划版本，线性时间复杂度
func Fib2(n int) int {
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

//logn时间复杂度，高效

func Fib3(n int) int {
	if n < 2 {
		return n
	}

	m := Fib3Part(n)
	return m.a1
}
func Fib3Part(n int) *SecondOrderMatrix {

	m := SecondOrderMatrix{
		a1: 1,
		a2: 1,
		b1: 1,
		b2: 0,
	}

	if n == 1 {
		return &m
	}

	if n%2 == 0 {
		t := SecondOrderMatrixMutl(Fib3Part(n/2), Fib3Part(n/2))
		return t
	} else {
		t := SecondOrderMatrixMutl(SecondOrderMatrixMutl(Fib3Part(n/2), Fib3Part(n/2)), &m)
		return t
	}
}

type SecondOrderMatrix struct {
	a1, a2 int
	b1, b2 int
}

//二阶矩阵的乘法
func SecondOrderMatrixMutl(s1, s2 *SecondOrderMatrix) *SecondOrderMatrix {
	r := &SecondOrderMatrix{}
	r.a1 = s1.a1*s2.a1 + s1.b1*s2.a2
	r.a2 = s1.a1*s2.a2 + s1.a2*s2.b2
	r.b1 = s1.b1*s2.a1 + s1.b2*s2.b1
	r.b2 = s1.b1*s2.a2 + s1.b2*s2.b2
	return r
}
