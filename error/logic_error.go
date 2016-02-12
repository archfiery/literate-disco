package error

import (
	"fmt"
)

// Index out of range error
// It is usually thrown for a slice
type OutOfRangeError struct {}

// Returns a string describing the error
func (err OutOfRangeError) Error() string {
	return fmt.Sprintf("Index Out Of Range")
}

// Types of object do not match error
// T1 and T2 are the names of two types respectively
type TypeNotMatchError struct {
	T1   string
	T2   string
}

// Returns a string describing the error
func (err TypeNotMatchError) Error() string {
	return fmt.Sprintf("Type %v and type %v do not match", err.T1, err.T2)
}
