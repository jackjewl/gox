package queue

import (
	"fmt"
	"testing"
)

func TestNewSequenceQueue(t *testing.T) {
	var q Queue[int] = NewSequenceQueue[int]()

	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	fmt.Printf("q.Size(): %v\n", q.Size())
	fmt.Printf("q.Empty(): %v\n", q.Empty())
	fmt.Printf("q.DeQueue(): %v\n", q.DeQueue())
	fmt.Printf("q.Front(): %v\n", q.Front())

}
