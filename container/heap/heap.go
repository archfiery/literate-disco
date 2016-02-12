package heap

import(
	"github.com/archfiery/literate-disco/container"
)

// A default Heap struct
// It is a binary heap, using a slice to store data
// A comparison function must be supplied for swapping condition
type Heap struct {
	data []interface{}
	comp container.CompFunc
}

// Return the item by index
func (h Heap) At(i int) interface{} {
	return h.data[i]
}

// Swap the element of internal array by index
func (h Heap) Swap(i int, j int) {
	swap(h.data, i, j)
}

// For all the items in the data array
// Run heapify operations
func (h Heap) BuildHeap() {
	heapifyAll(h.data, h.comp)
}

// Clear all data for this heap
func (h *Heap) Clear() {
	h.data = h.data[:0]
}

// Return true if the heap is empty
// false otherwise
func (h Heap) Empty() bool {
	return (h.Size() == 0)
}

// Return the number items for this heap
func (h Heap) Size() int {
	return len(h.data)
}

// Sort all the element by HeapSort
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

// Swap element for a slice by index
func swap(A []interface{}, a int, b int) {
	A[a], A[b] = A[b], A[a]
}

// Return the index of left child
func left(i int) int {
	return i*2 + 1
}

// Return the index of right child
func right(i int) int {
	return i*2 + 2
}

// Return the index of parent
func parent(i int) int {
	return i / 2
}

// Run heapify function on all items for an array
func heapifyAll(A []interface{}, f container.CompFunc) {
	size := len(A)
	for i := size / 2; i >= 0; i-- {
		heapify(A, i, f)
	}
}

// Recursively run heapify operation
// To maintain the heap property
func heapify(A []interface{}, i int, f container.CompFunc) {
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
