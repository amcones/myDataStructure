package priorityQueue

import (
	"errors"
	"golang.org/x/exp/constraints"
)

// PriorityQueue is a data structure based on heap. It can insert and auto sort the values in O(log n).
type PriorityQueue[T constraints.Ordered] struct {
	array []T
	N     int
	op    bool
}

func NewPriorityQueue[T constraints.Ordered](less bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		array: make([]T, 1),
		N:     0,
		op:    less,
	}
}

func (q *PriorityQueue[T]) parent(root int) int {
	return root / 2
}

func (q *PriorityQueue[T]) less(a, b T) bool {
	if q.op {
		return a < b
	} else {
		return a > b
	}
}

func (q *PriorityQueue[T]) left(root int) int {
	return 2 * root
}

func (q *PriorityQueue[T]) right(root int) int {
	return 2*root + 1
}

func (q *PriorityQueue[T]) swim(k int) {
	for q.parent(k) != 0 && q.less(q.array[k], q.array[q.parent(k)]) {
		q.swap(k, q.parent(k))
		k >>= 1
	}
}

func (q *PriorityQueue[T]) sink(k int) {
	t := k
	if q.left(k) <= q.N && q.less(q.array[q.left(k)], q.array[t]) {
		t = q.left(k)
	}
	if q.right(k) <= q.N && q.less(q.array[q.right(k)], q.array[t]) {
		t = q.right(k)
	}
	if k != t {
		q.swap(k, t)
		q.sink(t)
	}
}

func (q *PriorityQueue[T]) swap(i, j int) {
	temp := q.array[i]
	q.array[i] = q.array[j]
	q.array[j] = temp
}

// PushBack push value into priority queue
func (q *PriorityQueue[T]) PushBack(e T) {
	q.N++
	q.array = append(q.array, e)
	q.swim(q.N)
}

// Top return the maximum value of the priority queue
func (q *PriorityQueue[T]) Top() (T, error) {
	var zeroT T
	if q.N == 0 {
		return zeroT, errors.New("the priority queue is empty")
	}
	return q.array[1], nil
}

// PopBack pop the maximum value of the priority queue
func (q *PriorityQueue[T]) PopBack() error {
	if q.N == 0 {
		return errors.New("the priority queue is empty")
	}
	q.swap(1, q.N)
	q.N--
	q.sink(1)
	return nil
}

// Empty return the priority queue is empty or not
func (q *PriorityQueue[T]) Empty() bool {
	return q.N == 0
}
