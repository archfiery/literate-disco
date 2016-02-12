package container
import "github.com/archfiery/literate-disco/container/error"

// A comparison function type
type CompFunc func(a interface{}, b interface{}) (bool, *error.TypeNotMatchError)
