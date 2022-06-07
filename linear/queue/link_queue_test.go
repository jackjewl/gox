package queue

import (
	"fmt"
	"testing"
)

func TestNewLinkQueue(t *testing.T) {
	lq := NewLinkQueue[int]()
	fmt.Printf("lq.Empty(): %v\n", lq.Empty())
	lq.EnQueue(1)
	lq.EnQueue(2)
	lq.EnQueue(3)
	fmt.Printf("lq.Elems: %v\n", lq.Elems)
	fmt.Printf("lq.DeQueue(): %v\n", lq.DeQueue())
	fmt.Printf("lq.Elems: %v\n", lq.Elems)
	fmt.Printf("lq.Front(): %v\n", lq.Front())

	var a Queue[int] = NewLinkQueue[int]()
	a.EnQueue(1)
	a.EnQueue(2)
	a.EnQueue(3)
	fmt.Printf("a.DeQueue(): %v\n", a.DeQueue())
	fmt.Printf("a.Front(): %v\n", a.Front())
	fmt.Printf("a.Empty(): %v\n", a.Empty())
	fmt.Printf("a.Size(): %v\n", a.Size())
}
