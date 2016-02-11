package vector

import (
	"github.com/archfiery/literate-disco/container"
	"github.com/archfiery/literate-disco/container/error"
	"time"
)

const (
	INIT_CAP = 1024
)

type Vector struct {
	data []interface{}
	comp container.CompFunc
}

func MakeVector(comp container.CompFunc) Vector {
	slice := make([]interface{}, 0, INIT_CAP)
	v := Vector{slice, comp}
	return v
}

func (v *Vector) Clear() {
	v.data = make([]interface{}, 0, INIT_CAP)
}

func (v Vector) Empty() bool {
	return (v.Size() == 0)
}

func (v Vector) Size() int {
	return len(v.data)
}

func (v Vector) MaxSize() int {
	return cap(v.data)
}

func (v *Vector) PushBack(a interface{}) {
	if (moreThanHalf(v.data)) {
		v.data = doubleSlice(&v.data)
	}
	v.data = append(v.data, a)
}

func (v Vector) Front() interface{} {
	return v.data[0]
}

func (v Vector) At(i int) (interface{}, *error.OutOfRangeError) {
	if i < 0 || i > len(v.data) - 1 {
		return 0, &error.OutOfRangeError{time.Now()}
	}
	return v.data[i], nil
}

//================
//helper functions
//================

func moreThanHalf(A []interface{}) bool {
	if (len(A) <  cap(A) / 2) {
		return false
	}
	return true
}

func doubleSlice(A *[]interface{}) []interface{} {
	B := make([]interface{}, len(*A), 2 * cap(*A))
	copy(B, *A)
	return B
}
