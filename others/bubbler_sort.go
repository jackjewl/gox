package others

import "gox/utils"

//冒泡排序，这里对应证明程序正确性，使用了循环不变性和单调性质
func BubbleSort1[T any](elems []T, less func(T, T) bool) {

	sorted := false
	for !sorted {
		sorted = true
		for i := 1; i < len(elems); i++ {
			if less(elems[i], elems[i-1]) {
				{
					utils.Swap(&elems[i-1], &elems[i])
				}
				sorted = false
			}
		}
		elems = elems[:len(elems)-1]
	}
}
