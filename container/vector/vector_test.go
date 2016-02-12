package vector

import (
	"fmt"
	"testing"
	"github.com/archfiery/literate-disco/container/error"
	"reflect"
)

// Test moreThanHalf() function in vector.go
func TestMoreThanHalf(t *testing.T) {
	fmt.Println("TestMoreThanHalf")
	A := make([]interface{}, 0, 4)
	if moreThanHalf(A) != false {
		t.Fatal("len(A) is ", len(A), " and cap(A) is ", cap(A))
	}
	for i := 0; i < 3; i++ {
		A = append(A, i)
	}
	if moreThanHalf(A) != true {
		t.Fatal("len(A) is ", len(A), " and cap(A) is ", cap(A))
	}
}

// Test doubleSlice() function in vector.go
func TestDoubleSlice(t *testing.T) {
	fmt.Println("TestDoubleSlice")
	prevCap := 10
	A := make([]interface{}, 0, prevCap)
	B := doubleSlice(&A)
	if cap(B) != prevCap*2 {
		t.Fatal("cap are not doubled, expected ", prevCap*2, ", but get ", cap(B))
	}
}

// A lessThan function with type matching
// It throws TypeNotMatch error when a and b are not of the same type
func lessThanVectorTest(a interface{}, b interface{}) (bool, *error.TypeNotMatchError) {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		st1, st2 := reflect.TypeOf(a).String(), reflect.TypeOf(b).String()
		return false, &error.TypeNotMatchError{st1, st2}
	}
	switch a := a.(type) {
	case int:
		if a < b.(int) {
			return true, nil
		}
	case byte:
		if a < b.(byte) {
			return true, nil
		}
	case string:
		if a < b.(string) {
			return true, nil
		}
	}
	return false, nil
}

// Test PushBack() function in vector.go
// It tests int, byte and string
func TestPushBack(t *testing.T) {
	fmt.Println("TestPushBack")
	f := lessThanVectorTest
	vec := MakeVector(f)
	vec.PushBack(100)
	vec.PushBack(100)
	vec.PushBack(1)

	val, err := vec.At(2)
	if err != nil || val != 1 {
		t.Fatal("The push does not work properly!")
	}

	vec2 := MakeVector(f)
	vec2.PushBack('A' + 35)
	val, err = vec2.At(0)
	if err != nil || val != 'A'+35 {
		t.Fatal("The push does not work properly!")
	}

	vec3 := MakeVector(f)
	vec3.PushBack("literal")
	val, err = vec3.At(0)
	if err != nil || val != "literal" {
		t.Fatal("The push does not work properly!")
	}

	if vec.Size() != 3 {
		t.Fatal("The length of underlying array does not increment properly")
	}

	if vec.MaxSize() != INIT_CAP {
		fmt.Println("excpected: ", INIT_CAP, " actual: ", vec.MaxSize())
		t.Fatal("The maximal allowed size should not exceed initial cap in this case")
	}
}
