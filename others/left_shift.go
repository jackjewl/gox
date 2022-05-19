package others

func LeftShift[T any](elems []T, distance int) {

	Reverse[T](elems[:distance])
	Reverse[T](elems[distance:])
	Reverse[T](elems)

}

func RightShift[T any](elems []T, distance int) {
	Reverse[T](elems[len(elems)-distance:])
	Reverse[T](elems[len(elems)-distance:])
	Reverse[T](elems)
}
