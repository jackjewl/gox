package others

//数组元素整体左移distance位
func LeftShift[T any](elems []T, distance int) {
	Reverse1[T](elems[:distance])
	Reverse1[T](elems[distance:])
	Reverse1[T](elems)

}

//数组元素整体右移distance位
func RightShift[T any](elems []T, distance int) {
	Reverse1[T](elems[len(elems)-distance:])
	Reverse1[T](elems[:len(elems)-distance])
	Reverse1[T](elems)
}
