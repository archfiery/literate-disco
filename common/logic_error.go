// Package error provides error type
package common

import (
	"fmt"
)

// Index out of range error
// It is usually thrown for a slice
type OutOfRangeError struct {
	Msg string
}

// Returns a string describing the error
func (err OutOfRangeError) Error() string {
	return fmt.Sprintf("Index Out Of Range")
}

// Types of object do not match error
// T1 and T2 are the names of two types respectively
type TypeNotMatchError struct {
	T1 string
	T2 string
}

// Returns a string describing the error
func (err TypeNotMatchError) Error() string {
	return fmt.Sprintf("Type %v and type %v do not match", err.T1, err.T2)
}

// Assertion Error
// It can be thrown by any function
type AssertionError struct{}

// Returns a string describing the error
func (err AssertionError) Error() string {
	return fmt.Sprintf("Assertion error")
}

// NoSuchElementError
// It is usually thrown when trying to access some element that does not exist in the container
// For example: access the first element from a linked list when the list is empty
type NoSuchElementError struct{}

// Returns a string describing the error
func (err NoSuchElementError) Error() string {
	return fmt.Sprintf("No Such Element error")
}
