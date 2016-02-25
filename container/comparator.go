package container

import "github.com/archfiery/literate-disco/error"

// A comparison function type
type CompFunc func(a interface{}, b interface{}) (bool, error.Error)

// A equal function type
type EqualFunc func(a interface{}, b interface{}) (bool, error.Error)
