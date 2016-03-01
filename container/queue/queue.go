package queue

import "github.com/archfiery/literate-disco/common"

const (
	INIT_CAP = 1024
)

// A simple queue implementation without concurrent support
// It uses an array as underlying data structure, and a slice to describe data
// The queue supports arbitrary number of elements, it will resize when nearly overflow
type Queue struct {
	data []interface{}
}

func MakeQueue() Queue {
	s := make([]interface{}, 0, INIT_CAP)
	q := Queue{s}

	return q
}

//==========
// Capacity
//==========

// Returns true if the queue is empty, false otherwise
func (q Queue) Empty() bool {
	return (q.Size() == 0)
}

// Returns the number of items in the queue
func (q Queue) Size() int {
	return len(q.data)
}

// Returns the maximum allowed items for the queue
// It is the capacity of underlying slice
func (q Queue) Capacity() int {
	return cap(q.data)
}

//===========
// Modifiers
//===========

func (q *Queue) Clear() {
	q.data = make([]interface{}, 0, INIT_CAP)
}

func (q *Queue) Push(e interface{}) {
	// double the array if necessary
	if moreThanHalf(q.data) {
		q.data = doubleSlice(&q.data)
	}
	q.data = append(q.data, e)
}

func (q *Queue) Pop() error {
	if q.Size() <= 0 {
		return common.OutOfRangeError{}
	}
	if lessThanQuarter(q.data) && cap(q.data) > INIT_CAP {
		q.data = halveSlice(&q.data)
	}
	_, q.data = q.data[0], q.data[1:]

	return nil
}

func (q *Queue) Enqueue(e interface{}) {
	q.Push(e)
}

func (q *Queue) Dequeue() (interface{}, error) {
	e, err := q.Front()
	if err != nil {
		return -1, err
	}
	q.Pop()
	return e, nil
}

//================
// Element Access
//================
func (q Queue) Front() (interface{}, error) {
	if q.Size() <= 0 {
		return -1, common.OutOfRangeError{}
	}
	return q.data[0], nil
}

func (q Queue) Back() (interface{}, error) {
	if q.Size() <= 0 {
		return -1, common.OutOfRangeError{}
	}
	return q.data[q.Size()-1], nil
}

//================
//helper functions
//================

// Returns true if len(data) >= cap(data) / 2
func moreThanHalf(A []interface{}) bool {
	if len(A) < cap(A)/2 {
		return false
	}
	return true
}

// Returns true if len(data) < cap(data) / 4
func lessThanQuarter(A []interface{}) bool {
	if len(A) < cap(A)/4 {
		return true
	}
	return false
}

// Returns a new slice with doubled capacity of the original one
// Copies the items from original slice to the new one
func doubleSlice(A *[]interface{}) []interface{} {
	return alterCapacity(A, cap(*A)*2)
}

// Returns a new slice with half capacity of the original one
// Copies the items from original slice to the new one
func halveSlice(A *[]interface{}) []interface{} {
	return alterCapacity(A, cap(*A)/2)
}

// Returns a new slice with size n
// Copies the items from original slice to the new one
func alterCapacity(A *[]interface{}, n int) []interface{} {
	B := make([]interface{}, len(*A), n)
	copy(B, *A)
	return B
}
