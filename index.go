package list

import (
	"github.com/intdxdt/deque"
	"github.com/intdxdt/math"
)

const N = 32

// List - Type
type List[T any] struct {
	*deque.Deque[T]
}

// NewList - constructor the list type
func NewList[T any](initSize ...int) *List[T] {
	var iSize = N
	if len(initSize) > 0 {
		iSize = math.Max(1, initSize[0])
	}
	return &List[T]{deque.NewDeque[T](iSize)}
}

// Clear - drops entire list - similar to empty the list
func (lst *List[T]) Clear() {
	lst.Deque.Clear()
}

// Len - computes the length of list
func (lst *List[T]) Len() int {
	return lst.Deque.Len()
}

// First - gets first item in the list
func (lst *List[T]) First() T {
	return lst.Get(0)
}

// Last - gets last item in the list
func (lst *List[T]) Last() T {
	return lst.Get(-1)
}

// Append - adds to end of list
func (lst *List[T]) Append(data T) *List[T] {
	lst.Deque.Append(data)
	return lst
}

// AppendLeft - adds item at start of the list
func (lst *List[T]) AppendLeft(data T) *List[T] {
	lst.Deque.AppendLeft(data)
	return lst
}

// Pop - removes item from the last of the list
func (lst *List[T]) Pop() T {
	return lst.Deque.Pop()
}

// PopLeft - removes item from front of list
func (lst *List[T]) PopLeft() T {
	return lst.Deque.PopLeft()
}

// At finds the item at given at index n - O(1)
// Supports negative indexing
func (lst *List[T]) At(n int) T {
	if n < 0 {
		n = lst.Len() + n
	}
	return lst.Deque.Get(n)
}

// Each - iterates over items in the list
func (lst *List[T]) Each(fn func(T, int) bool) {
	lst.ForEach(func(v T, index int) bool {
		fn(v, index)
		return true
	})
}

// Stringer method
func (lst *List[T]) String() string {
	return lst.Deque.String()
}
