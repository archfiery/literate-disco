package vector

import (
	"fmt"
	"github.com/archfiery/literate-disco/error"
	"reflect"
	"testing"
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

// Test lessThanQuater() function in vector.go
func TestLessThanQuarter(t *testing.T) {
	fmt.Println("TestLessThanQuarter")
	A := make([]interface{}, 0, 4)
	if lessThanQuarter(A) != true {
		t.Fatal("len(A) is ", len(A), " and cap(A) is ", cap(A))
	}
	for i := 0; i < 3; i++ {
		A = append(A, i)
	}
	if lessThanQuarter(A) != false {
		t.Fatal("len(A) is ", len(A), " and cap(A) is ", cap(A))
	}
}

// Test doubleSlice() function in vector.go
func TestHalveSlice(t *testing.T) {
	fmt.Println("TestHalveSlice")
	prevCap := 20
	A := make([]interface{}, 0, prevCap)
	B := halveSlice(&A)
	if cap(B) != prevCap/2 {
		t.Fatal("cap are not doubled, expected ", prevCap/2, ", but get ", cap(B))
	}
}

// A lessThan function with type matching
// It throws TypeNotMatch error when a and b are not of the same type
func lessThanVectorTest(a interface{}, b interface{}) (bool, error.Error) {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		st1, st2 := reflect.TypeOf(a).String(), reflect.TypeOf(b).String()
		return false, error.TypeNotMatchError{st1, st2}
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

	if vec.Capacity() != INIT_CAP {
		fmt.Println("excpected: ", INIT_CAP, " actual: ", vec.Capacity())
		t.Fatal("The maximal allowed size should not exceed initial cap in this case")
	}
}

// Test PopBack() function in vector.go
func TestPopBack(t *testing.T) {
	fmt.Println("TestPushBack")
	f := lessThanVectorTest
	vec := MakeVector(f)
	vec.PushBack(100)
	vec.PushBack(100)
	vec.PushBack(1)

	val, err := vec.At(2)
	if err != nil || val != 1 || vec.Size() != 3 {
		t.Fatal("The items are not pushed to vector properly")
	}

	err = vec.PopBack()
	if err != nil {
		t.Fatal(err.Error())
	}

	val, err = vec.Back()
	if err != nil || val != 100 || vec.Size() != 2 {
		fmt.Printf("vec.At(1): %d; size: %d\n", val, vec.Size())
		t.Fatal("PopBack() does not function properly")
	}

	for i := 0; i < 10000; i++ {
		vec.PushBack(i)
	}
	if vec.Size() != 10002 {
		t.Fatal("vector is expected to have size of 10002, but it has size of ", vec.Size())
	}
	for i := 0; i < 6000; i++ {
		vec.PopBack()
	}
	if vec.Capacity() != 8192 {
		t.Fatal("vector is expected to have capacity of 8192, but it has capacity of ", vec.Capacity())
	}
	if vec.Size() < vec.Capacity()/4 {
		t.Fatal("capcity are not fairly bounded")
	}
}

// Test Insert() function in vector.go
func TestInsert(t *testing.T) {
	fmt.Println("TestInsert")
	f := lessThanVectorTest
	vec := MakeVector(f)
	vec.PushBack(100)
	vec.PushBack(100)
	vec.PushBack(1)

	// Test insert at 0
	err := vec.Insert(0, 5)
	if err != nil {
		t.Fatal(err.Error())
	}
	val, err := vec.At(0)
	if err != nil {
		t.Fatal(err.Error())
	}
	if val != 5 {
		fmt.Println(vec.data)
		t.Fatal("Wrong element")
	}
	if vec.Size() != 4 {
		t.Fatal("Size is incorrect")
	}

	vec.Insert(2, 15)
	val, _ = vec.At(2)
	if val != 15 {
		fmt.Println(vec.data)
		t.Fatal("Wrong element")
	}
	if vec.Size() != 5 {
		t.Fatal("Size is incorrect")
	}

	vec.Insert(5, 25)
	val, _ = vec.At(5)
	if val != 25 {
		fmt.Println(vec.data)
		t.Fatal("Wrong element")
	}
	if vec.Size() != 6 {
		t.Fatal("Size is incorrect")
	}
	fmt.Println(vec.data)
}
