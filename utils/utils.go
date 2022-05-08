package utils

func Swap(a *interface{}, b *interface{}) {
	*a, *b = *b, *a
}
