package vector

const (
	INIT_CAP = 1024
)

type Vector struct {
	data []interface{}
}

func (v *Vector) Init() {
	v.data = make([]interface{}, 0, INIT_CAP)
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
