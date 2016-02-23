package sort

import (
	"github.com/archfiery/literate-disco/container"
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

func Leq(a interface{}, b interface{}) (bool, error.Error) {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		st1 := reflect.TypeOf(a).String()
		st2 := reflect.TypeOf(b).String()
		return false, &error.TypeNotMatchError{st1, st2}

	}
	return a.(int) <= b.(int), nil
}

func Geq(a interface{}, b interface{}) (bool, error.Error) {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		st1 := reflect.TypeOf(a).String()
		st2 := reflect.TypeOf(b).String()
		return false, &error.TypeNotMatchError{st1, st2}
	}
	return a.(int) >= b.(int), nil
}

func IsSorted(A []interface{}, f container.CompFunc) bool {
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
