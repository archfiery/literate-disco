package vector

import (
	"github.com/archfiery/literate-disco/container"
	"github.com/archfiery/literate-disco/container/error"
)

const (
	// initial capacity of the vector
	INIT_CAP = 1024
)

// Type vector, using an array as underlying data structure
// and a slice to describe the array
// It also needs a comparison function that throws TypeNotMatch error
type Vector struct {
	data []interface{}
	comp container.CompFunc
}

// Return a new vector with initialised capacity
func MakeVector(comp container.CompFunc) Vector {
	slice := make([]interface{}, 0, INIT_CAP)
	v := Vector{slice, comp}
	return v
}

// Clear all data in the data
func (v *Vector) Clear() {
	v.data = make([]interface{}, 0, INIT_CAP)
}

// Return true if the vector is empty, false otherwise
func (v Vector) Empty() bool {
	return (v.Size() == 0)
}

// Return the number of items in the vector
func (v Vector) Size() int {
	return len(v.data)
}

// Return the maximum allowed items for the vector
// It is the capacity of underlying slice
func (v Vector) MaxSize() int {
	return cap(v.data)
}

// Pushes an item to the back of the vector
func (v *Vector) PushBack(a interface{}) {
	if moreThanHalf(v.data) {
		v.data = doubleSlice(&v.data)
	}
	v.data = append(v.data, a)
}

// Return the first element in the vector
func (v Vector) Front() (interface{}, *error.OutOfRangeError) {
	if (v.Size() <= 0) {
		return 0, &error.OutOfRangeError{}
	}
	return v.data[0], nil
}

// Return the last element in the vector
func (v Vector) Back() (interface{}, *error.OutOfRangeError) {
	if (v.Size() <= 0) {
		return 0, &error.OutOfRangeError{}
	}
	return v.data[v.Size() - 1], nil
}

// Return the item at index
func (v Vector) At(i int) (interface{}, *error.OutOfRangeError) {
	if i < 0 || i > v.Size()-1 {
		return 0, &error.OutOfRangeError{}
	}
	return v.data[i], nil
}

//================
//helper functions
//================

// Return true if the len(data) >= cap(data)
func moreThanHalf(A []interface{}) bool {
	if len(A) < cap(A)/2 {
		return false
	}
	return true
}

// Return a new slice with doubled capacity of the original one
// Copy the item from original slice to the new one
func doubleSlice(A *[]interface{}) []interface{} {
	B := make([]interface{}, len(*A), 2*cap(*A))
	copy(B, *A)
	return B
}
