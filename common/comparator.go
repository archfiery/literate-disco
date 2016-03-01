// Package common provides several function types that are commonly used
package common

// A comparison function type
type CompFunc func(a interface{}, b interface{}) (bool, error)

// A equal function type
type EqualFunc func(a interface{}, b interface{}) (bool, error)
