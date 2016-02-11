package error
import (
	"time"
	"fmt"
)

type OutOfRangeError struct {
	When time.Time
}

func (err OutOfRangeError) Error() string {
	return fmt.Sprintf("%v: Index Out Of Range", err.When)
}