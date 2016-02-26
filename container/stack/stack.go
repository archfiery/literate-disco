package stack

import (
	"github.com/archfiery/literate-disco/container"
	"github.com/archfiery/literate-disco/container/vector"
	"github.com/archfiery/literate-disco/error"
)

// Type stack, using a vector as underlying data structure
// It uses Adapter pattern that wraps up the vector
type Stack struct {
	vec  vector.Vector
	comp container.CompFunc
}

// Returns a new stack with a default vector
func MakeStack(comp container.CompFunc) Stack {
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
func (s Stack) Top() (interface{}, error.Error) {
	if s.Size() == 0 {
		return -1, error.OutOfRangeError{}
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
// Returns error if the stack is already empty
func (s *Stack) Pop() error.Error {
	if s.Size() == 0 {
		return error.OutOfRangeError{}
	}
	s.vec.PopBack()
	return nil
}
