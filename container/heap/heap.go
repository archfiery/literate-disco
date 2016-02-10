package heap

import ()

// A default Heap struct
// It is a binary heap, using a slice to store data
// A comparison function must be supplied for swapping condition
type Heap struct {
	data []interface{}
	comp func(a interface{}, b interface{}) bool
}

// Returns the item by a given index
func (h Heap) At(i int) interface{} {
	return h.data[i]
}

// For all the items in the data array
// Run heapify operations
func (h Heap) BuildHeap() {
	for i := h.Size() / 2; i >= 0; i-- {
		h.heapify(i)
	}
}

// Clear all data for this heap
func (h *Heap) Clear() {
	h.data = h.data[:0]
}

// Returns true if the heap is empty
// false otherwise
func (h Heap) Empty() bool {
	return (h.Size() == 0)
}

// Returns the number items for this heap
func (h Heap) Size() int {
	return len(h.data)
}

// ========================
// helper functions
// ========================

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
	return i / 2
}

// Recursively run heapify operation
// To maintain the heap property
func (h *Heap) heapify(i int) {
	l := left(i)
	r := right(i)
	need := i
	if l < h.Size() && h.comp(h.data[l], h.data[i]) {
		need = l
	} else {
		need = i
	}
	if r < h.Size() && h.comp(h.data[r], h.data[need]) {
		need = r
	}
	if need != i {
		h.data[i], h.data[need] = h.data[need], h.data[i]
		h.heapify(need)
	}
}
