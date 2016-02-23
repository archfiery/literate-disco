package sort

import (
	"fmt"
	"testing"
	test_util "github.com/archfiery/literate-disco/test"
)

// Test Partitioning function for quicksort
// The position of pivot only works for the non-random one
// As we use random method in partitioning, the test does not check the pivot any more
func TestQuicksortPartition(t *testing.T) {
	fmt.Println("TestQuicksortPartition")
	dataSlice := []int{27, 17, 3, 16, 13}
	array := make([]interface{}, len(dataSlice))
	for i, v := range dataSlice {
		array[i] = v
	}

	_, err := quicksortPartition(array, 0, 4, test_util.Leq)
	if err != nil {
		t.Fatal("quick sort partition error")
	}
	//if val != 1 {
	//	t.Fatal("pivot position is incorrect")
	//}
}

// Test if the Quicksort works
func TestQuicksort(t *testing.T) {
	fmt.Println("TestBuildHeap")
	dataSlice := []int{27, 17, 3, 16, 13, 10, 1, 5, 7, 12, 4, 8, 9, 0}
	array := make([]interface{}, len(dataSlice))
	for i, v := range dataSlice {
		array[i] = v
	}
	Quicksort(array, test_util.Leq)
	fmt.Println(array)
	if test_util.IsSorted(array, test_util.Leq) == false {
		t.Fatal("not sorted with leq function")
	}
	Quicksort(array, test_util.Geq)
	fmt.Println(array)
	if test_util.IsSorted(array, test_util.Geq) == false {
		t.Fatal("not sorted with geq function")
	}

}
