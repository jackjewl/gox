package stack

import (
	"log"
	"testing"
)

func TestNewLinkedStack(t *testing.T) {
	// ls := NewLinkedStack[int]()
	// log.Println(ls.elems)
	// log.Println(ls.Empty())
	// ls.Push(1)
	// ls.Push(2)
	// ls.Push(3)
	// log.Println(ls.elems)
	// log.Println(ls.Size())
	// log.Println(ls.Empty())
	// log.Println(ls.Pop())
	// log.Println(ls.elems)
	// log.Println(ls.Top())

	var s Stack[int] = NewLinkedStack[int]()
	log.Print(s)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	log.Println(s.Empty())
	log.Println(s.Size())
	log.Println(s.Top())
	log.Println(s.Pop())
	log.Println(s.Size())
}
