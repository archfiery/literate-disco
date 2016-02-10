package heap

import ()

type Heap struct {
	data []interface{}
	comp func(a interface{}, b interface{}) bool
}

func (c Heap) At(i int) interface{} {
	return c.data[i]
}

func (c Heap) BuildHeap() {
	for i := c.Size() / 2; i >= 0; i-- {
		c.heapify(i)
	}
}

func (c *Heap) Clear() {
	c.data = make([]interface{}, 0)
}

func (c Heap) Empty() bool {
	return (c.Size() == 0)
}

func (c Heap) Size() int {
	return len(c.data)
}

func (c Heap) Heapify(i int) {
	c.heapify(i)
}

// ========================
// helper functions
// ========================

func left(i int) int {
	return i*2 + 1
}

func right(i int) int {
	return i*2 + 2
}

func parent(i int) int {
	return i / 2
}

func (c *Heap) heapify(i int) {
	l := left(i)
	r := right(i)
	need := i
	if l < c.Size() && c.comp(c.data[l], c.data[i]) {
		need = l
	} else {
		need = i
	}
	if r < c.Size() && c.comp(c.data[r], c.data[need]) {
		need = r
	}
	if need != i {
		c.data[i], c.data[need] = c.data[need], c.data[i]
		c.Heapify(need)
	}
}
