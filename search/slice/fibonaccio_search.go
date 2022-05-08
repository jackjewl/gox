package slice

import (
	"gox/fib"
)

func FibonacciSearch(target interface{}, elems []interface{}, compare func(interface{}, interface{}) int) int {
	return FibonacciSearchSection(target, elems, compare, 0, len(elems))
}

func FibonacciSearchSection(target interface{}, elems []interface{}, compare func(interface{}, interface{}) int, low, high int) int {

	fibonacci := fib.Fib{}
	fibonacci.Generate(high - low)
	var mid int

	for low < high {
		for high-low > fibonacci.Get() {
			fibonacci.Pre()
		}

		mid = low + fibonacci.Get() - 1
		if compare(target, elems[mid]) < 0 {
			high = mid
		} else if compare(target, elems[mid]) > 0 {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
