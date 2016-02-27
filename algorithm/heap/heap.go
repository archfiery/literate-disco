// Package heap implements the heap data structure
package heap

import (
	c "github.com/archfiery/literate-disco/common"
)

// A default Heap struct
// It is a binary heap, using a slice to store data
// A comparison function must be supplied for swapping condition
type Heap struct {
	data []interface{}
	comp c.CompFunc
}

// Returns a new binary heap by a given data and comparison function
func MakeHeap(data []interface{}, comp c.CompFunc) Heap {
	return Heap{data, comp}
}

// Returns the element by index
func (h Heap) At(i int) interface{} {
	return h.data[i]
}

// Swaps the element on internal array by index
func (h Heap) Swap(i int, j int) {
	swap(h.data, i, j)
}

// For all the elements in the data array
// Runs heapify operations
func (h Heap) BuildHeap() {
	heapifyAll(h.data, h.comp)
}

// Clears all elements in the heap
func (h *Heap) Clear() {
	h.data = h.data[:0]
}

// Returns true if the heap is empty
// false otherwise
func (h Heap) Empty() bool {
	return (h.Size() == 0)
}

// Returns the number of elements for this heap
func (h Heap) Size() int {
	return len(h.data)
}

// Sorts all the elements by HeapSort
func (h Heap) HeapSort() {
	h.BuildHeap()
	l := len(h.data)
	for i := l - 1; i > 0; i-- {
		swap(h.data, 0, i)
		h.data = h.data[:len(h.data)-1]
		heapify(h.data, 0, h.comp)
	}
}

// ========================
// helper functions
// ========================

// Swaps element for a slice by index
func swap(A []interface{}, a int, b int) {
	A[a], A[b] = A[b], A[a]
}

// Returns the index of left child
func left(i int) int {
	return i*2 + 1
}

// Returns the index of right child
func right(i int) int {
	return i*2 + 2
}

// Returns the index of parent
func parent(i int) int {
	return (i - 1) / 2
}

// Runs heapify function on all items for an array
func heapifyAll(A []interface{}, f c.CompFunc) {
	size := len(A)
	for i := size / 2; i >= 0; i-- {
		heapify(A, i, f)
	}
}

// Recursively runs the heapify operation in order to maintain the heap property
// at the subtree of rooted at i
func heapify(A []interface{}, i int, f c.CompFunc) {
	l := left(i)
	r := right(i)
	need := i
	size := len(A)
	if l < size {
		val, err := f(A[l], A[i])
		if err == nil && val == true {
			need = l
		}
	} else {
		need = i
	}
	if r < size {
		val, err := f(A[r], A[need])
		if err == nil && val == true {
			need = r
		}
	}
	if need != i {
		A[i], A[need] = A[need], A[i]
		heapify(A, need, f)
	}
}
