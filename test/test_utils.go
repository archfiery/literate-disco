// Package test provides commonly used functions for running test cases
package test

import (
	"github.com/archfiery/literate-disco/common"
	"github.com/archfiery/literate-disco/error"
	"reflect"
)

// Test int data only
func LessThan(a interface{}, b interface{}) (bool, error.Error) {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		st1 := reflect.TypeOf(a).String()
		st2 := reflect.TypeOf(b).String()
		return false, &error.TypeNotMatchError{st1, st2}

	}
	return a.(int) < b.(int), nil
}

// Test int data only
func MoreThan(a interface{}, b interface{}) (bool, error.Error) {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		st1 := reflect.TypeOf(a).String()
		st2 := reflect.TypeOf(b).String()
		return false, &error.TypeNotMatchError{st1, st2}
	}
	return a.(int) > b.(int), nil
}

// Less than or equal to
func Leq(a interface{}, b interface{}) (bool, error.Error) {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		st1 := reflect.TypeOf(a).String()
		st2 := reflect.TypeOf(b).String()
		return false, &error.TypeNotMatchError{st1, st2}

	}
	return a.(int) <= b.(int), nil
}

// Greater than or equal to
func Geq(a interface{}, b interface{}) (bool, error.Error) {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		st1 := reflect.TypeOf(a).String()
		st2 := reflect.TypeOf(b).String()
		return false, &error.TypeNotMatchError{st1, st2}
	}
	return a.(int) >= b.(int), nil
}

// Returns true if the slice of interface{} is sorted based on the comparision function
func IsSorted(A []interface{}, f common.CompFunc) bool {
	if len(A) < 2 {
		return true
	}
	for i := 0; i < len(A)-1; i++ {
		val, err := f(A[i], A[i+1])
		if err != nil || val == false {
			return false
		}
	}
	return true
}

// An int equal function
func IntEqual (a interface{}, b interface{}) (bool, error.Error) {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		st1 := reflect.TypeOf(a).String()
		st2 := reflect.TypeOf(b).String()
		return false, &error.TypeNotMatchError{st1, st2}
	}
	return a.(int) == b.(int), nil
}
