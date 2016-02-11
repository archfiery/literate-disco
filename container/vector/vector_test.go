package vector

import (
	"fmt"
	"testing"
)

func TestIsHalf(t *testing.T) {
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

func TestDoubleSlice(t *testing.T) {
	fmt.Println("TestDoubleSlice")
	prevCap := 10
	A := make([]interface{}, 0, prevCap)
	B := doubleSlice(&A)
	if cap(B) != prevCap * 2 {
		t.Fatal("cap are not doubled, expected ", prevCap * 2, ", but get ", cap(B))
	}
}
