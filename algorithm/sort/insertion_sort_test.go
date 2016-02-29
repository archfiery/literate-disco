package sort

import (
	"fmt"
	"github.com/archfiery/literate-disco/test"
	"testing"
)

func TestBinaryInsertionSort(t *testing.T) {
	fmt.Println("\nTestBinaryInsertionSort")
	dataSlice := []int{27, 17, 3, 16, 13, 10, 1, 5, 7, 12, 4, 8, 9, 0}
	array := make([]interface{}, len(dataSlice))
	for i, v := range dataSlice {
		array[i] = v
	}

	BinaryInsertionSort(array, len(array), test.LessThan)

	for i := 0; i < len(array)-1; i++ {
		if v, err := test.LessThan(array[i], array[i+1]); err == nil && v == false {
			t.Fatal("Array fails to sort in ascending order")
		}
	}

	for i, v := range dataSlice {
		array[i] = v
	}

	BinaryInsertionSort(array, len(array), test.MoreThan)

	for i := 0; i < len(array)-1; i++ {
		if v, err := test.MoreThan(array[i], array[i+1]); err == nil && v == false {
			t.Fatal("Array fails to sort in descending order")
		}
	}
}
