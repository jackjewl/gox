package fib

type Fib struct {
	pre  int
	succ int
}

func (this *Fib) Generate(n int) {

	this.pre = 0
	this.succ = 1
	for this.succ < n {
		this.Next()
	}

}

func (this *Fib) Pre() {

	this.pre = this.succ - this.pre
	this.succ = this.succ - this.pre
}

func (this *Fib) Next() {
	this.succ = this.succ + this.pre
	this.pre = this.succ - this.pre

}

func (this *Fib) Get() int {
	return this.pre
}
