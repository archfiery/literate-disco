package container
import "github.com/archfiery/literate-disco/error"

// A comparison function type
type CompFunc func(a interface{}, b interface{}) (bool, *error.TypeNotMatchError)
