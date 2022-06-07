package stack

import (
	"log"
	"testing"
)

func TestNewSequenceStack(t *testing.T) {
	var s Stack[int] = NewSequenceStack[int]()
	log.Println(s.Empty())
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
