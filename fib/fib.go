package fib

type Fib struct {
	pre  int
	succ int
}

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	pre, succ := 0, 1
	for i := 1; i < n; i++ {
		succ = pre + succ
		pre = succ - pre
	}
	return succ
}

func (this *Fib) Generate(n int) {

	this.pre = 0
	this.succ = 1
	for this.succ < n {
		this.Next()
	}

}

func (this *Fib) Pre() {

	this.pre = this.succ
	this.succ = this.succ - this.pre
}

func (this *Fib) Next() {
	this.succ = this.succ + this.pre
	this.pre = this.succ - this.pre

}

func (this *Fib) Get() int {
	return this.pre
}
