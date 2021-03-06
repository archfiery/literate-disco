// Package stack implements stack data structure
package stack

import (
	"github.com/archfiery/literate-disco/common"
	"github.com/archfiery/literate-disco/container/vector"
)

// Type stack, using a vector as underlying data structure
// It uses Adapter pattern that wraps up the vector
type Stack struct {
	vec  vector.Vector
	comp common.CompFunc
}

// Returns a new stack with a default vector
func MakeStack(comp common.CompFunc) Stack {
	v := vector.MakeVector(comp)
	s := Stack{v, comp}
	return s
}

// Returns true if the stack is empty, false otherwise
func (s Stack) Empty() bool {
	return s.vec.Size() == 0
}

// Returns the number of elements in the stack
func (s Stack) Size() int {
	return s.vec.Size()
}

// Returns the top element
// Returns OutOfRangeError if the stack is already empty
func (s Stack) Top() (interface{}, error) {
	if s.Size() == 0 {
		return -1, common.OutOfRangeError{}
	}
	val, err := s.vec.Back()
	if err != nil {
		return -1, err
	}
	return val, nil
}

// Inserts an element to the stack
func (s *Stack) Push(i interface{}) {
	s.vec.PushBack(i)
}

// Remove the top element, reduces the size of stack by 1
// Returns common.if the stack is already empty
func (s *Stack) Pop() error {
	if s.Size() == 0 {
		return common.OutOfRangeError{}
	}
	s.vec.PopBack()
	return nil
}
