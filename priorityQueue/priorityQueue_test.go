package priorityQueue

import (
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue[string](true)
	pq.PushBack("abc")
	pq.PushBack("abcd")
	pq.PushBack("bcd")
	pq.PushBack("aac")
	pq.PushBack("aab")
	pq.PushBack("abb")
	for !pq.Empty() {
		res, _ := pq.Top()
		fmt.Println(res)
		_ = pq.PopBack()
	}
}
