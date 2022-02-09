package list

import (
	"github.com/intdxdt/deque"
	"github.com/intdxdt/math"
)

const N = 32

//List - Type
type List struct {
	*deque.Deque
}

//NewList - constructor the list type
func NewList(initSize ...int) *List {
	var iSize = N
	if len(initSize) > 0 {
		iSize = math.MaxInt(1, initSize[0])
	}
	return &List{deque.NewDeque(iSize)}
}

//Drop - drops entire list - similar to empty the list
func (lst *List) Clear() {
	lst.Deque.Clear()
}

//Len - computes the length of list
func (lst *List) Len() int {
	return lst.Deque.Len()
}

//First - gets first item in the list
func (lst *List) First() interface{} {
	if lst.Len() > 0 {
		return lst.Get(0)
	}
	return nil
}

//Last - gets last item in the list
func (lst *List) Last() interface{} {
	if lst.Len() > 0 {
		return lst.Get(-1)
	}
	return nil
}

//Append - adds to end of list
func (lst *List) Append(data interface{}) *List {
	lst.Deque.Append(data)
	return lst
}

//AppendLeft - adds item at start of the list
func (lst *List) AppendLeft(data interface{}) *List {
	lst.Deque.AppendLeft(data)
	return lst
}

//Pop - removes item from the last of the list
func (lst *List) Pop() interface{} {
	if lst.Len() == 0 {
		return nil
	}
	return lst.Deque.Pop()
}

//PopLeft - removes item from front of list
func (lst *List) PopLeft() interface{} {
	if lst.Len() == 0 {
		return nil
	}
	return lst.Deque.PopLeft()
}

//At finds the item at given at index n - O(1)
//Supports negative indexing
func (lst *List) At(n int) interface{} {
	if n < 0 {
		n = lst.Len() + n
	}
	if n >= 0 && n < lst.Len() {
		return lst.Deque.Get(n)
	}
	return nil
}

//Each - iterates over items in the list
func (lst *List) Each(fn func(interface{}, int) bool) {
	lst.ForEach(func(v interface{}, index int) bool {
		fn(v, index)
		return true
	})
}

//Stringer method
func (lst *List) String() string {
	return lst.Deque.String()
}
