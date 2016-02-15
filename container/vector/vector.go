package vector

import (
	"github.com/archfiery/literate-disco/container"
	"github.com/archfiery/literate-disco/error"
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

// Returns a new vector with initialised capacity
func MakeVector(comp container.CompFunc) Vector {
	slice := make([]interface{}, 0, INIT_CAP)
	v := Vector{slice, comp}
	return v
}

//==========
// Capacity
//==========


// Returns true if the vector is empty, false otherwise
func (v Vector) Empty() bool {
	return (v.Size() == 0)
}

// Returns the number of items in the vector
func (v Vector) Size() int {
	return len(v.data)
}

// Returns the maximum allowed items for the vector
// It is the capacity of underlying slice
func (v Vector) Capacity() int {
	return cap(v.data)
}


//==========
// Modifiers
//==========

// Clears all data in the data
func (v *Vector) Clear() {
	v.data = make([]interface{}, 0, INIT_CAP)
}

// Adds a new element at the end of the vector
func (v *Vector) PushBack(a interface{}) {
	if moreThanHalf(v.data) {
		v.data = doubleSlice(&v.data)
	}
	v.data = append(v.data, a)
}

// Removes the last element in the vector
func (v *Vector) PopBack() *error.OutOfRangeError {
	if (v.Size() <= 0) {
		return new(error.OutOfRangeError)
	}
	v.data = v.data[:v.Size() - 1]
	return nil
}

// Insert a single element to vector at position `index`
func (v *Vector) Insert(index int, a interface{}) *error.OutOfRangeError {
	if moreThanHalf(v.data) {
		v.data = doubleSlice(&v.data)
	}
	// when the index is invalid
	if index > v.Size() || index < 0{
		return new(error.OutOfRangeError)
	}
	// when the index is the last
	if index == v.Size() {
		v.PushBack(a)
		return nil
	}
	// other valid index
	first := v.data[:index]
	second := v.data[index:]
	v.data = make([]interface{}, 0, v.Capacity())
	v.data = append(v.data, first...)
	v.data = append(v.data, a)
	v.data = append(v.data, second...)

	return nil
}

//================
// Element Access
//================

// Returns the first element in the vector
func (v Vector) Front() (interface{}, *error.OutOfRangeError) {
	if v.Size() <= 0 {
		return 0, new(error.OutOfRangeError)
	}
	return v.data[0], nil
}

// Returns the last element in the vector
func (v Vector) Back() (interface{}, *error.OutOfRangeError) {
	if v.Size() <= 0 {
		return 0, &error.OutOfRangeError{}
	}
	return v.data[v.Size()-1], nil
}

// Returns the item at index
func (v Vector) At(i int) (interface{}, *error.OutOfRangeError) {
	if i < 0 || i > v.Size()-1 {
		return 0, new(error.OutOfRangeError)
	}
	return v.data[i], nil
}

//================
//helper functions
//================

// Returns true if the len(data) >= cap(data)
func moreThanHalf(A []interface{}) bool {
	if len(A) < cap(A)/2 {
		return false
	}
	return true
}

// Returns a new slice with doubled capacity of the original one
// Copies the item from original slice to the new one
func doubleSlice(A *[]interface{}) []interface{} {
	B := make([]interface{}, len(*A), 2*cap(*A))
	copy(B, *A)
	return B
}
