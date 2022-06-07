package set

type Set[T any] interface {
	Add(T) bool
	Remove() bool
	Exist(T) bool
	Empty() bool
	Size() int

	Union(Set[T]) *Set[T]
	Intersection(Set[T]) *Set[T]
	Complementary(universe Set[T]) *Set[T]
}
