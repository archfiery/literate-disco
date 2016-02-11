package heap

import (
	"fmt"
	"testing"
	"github.com/archfiery/literate-disco/container"
)

func LessThan(a interface{}, b interface{}) bool {
	return a.(int) < b.(int)
}

func MoreThan(a interface{}, b interface{}) bool {
	return a.(int) > b.(int)
}

// Test heapify() function in heap.go
func TestHeapify(t *testing.T) {
	fmt.Println("TestHeapify")
	// make dummy data
	dataSlice := []int{3, 1, 2}
	array := make([]interface{}, len(dataSlice))
	for i, v := range dataSlice {
		array[i] = v
	}
	// make heap
	heap := Heap{array, LessThan}
	// verify heap size
	if heap.Size() != 3 {
		t.Fatal("The size does not match")
	}
	// do heapify on the first item
	heapify(heap.data, 0, heap.comp)
	if heap.At(0) != 1 {
		t.Fatal("Heapify error")
	}
}

// Verify the heap property for an array according to its comparison function
func VerifyHeapProperty(heap Heap) bool {
	valid := true
	for i, _ := range heap.data {
		l := left(i)
		r := right(i)
		if l < heap.Size() {
			if !heap.comp(heap.At(i), heap.At(l)) {
				valid = false
			}
		}
		if r < heap.Size() {
			if !heap.comp(heap.At(i), heap.At(l)) {
				valid = false
			}
		}
	}
	return valid
}

// Test BuildHeap() function in heap.go
func TestBuildHeap(t *testing.T) {
	fmt.Println("TestBuildHeap")
	dataSlice := []int{27, 17, 3, 16, 13, 10, 1, 5, 7, 12, 4, 8, 9, 0}
	array := make([]interface{}, len(dataSlice))
	for i, v := range dataSlice {
		array[i] = v
	}
	// make heap
	heap := Heap{array, LessThan}
	heap.BuildHeap()

	// expected answer for test array
	expected := []int{0, 4, 1, 5, 12, 8, 3, 16, 7, 17, 13, 10, 9, 27}
	for i, v := range heap.data {
		if v != expected[i] {
			fmt.Println("Answer is not expected")
		}
	}
	// verify heap property according to comparison function
	if VerifyHeapProperty(heap) != true {
		fmt.Println("Heap property violated")
	}
}

// Test Size(), Clear() functions in heap.go
func TestBasicOps(t *testing.T) {
	fmt.Println("TestBasicOps")
	dataSlice := []int{27, 17, 3, 16, 13}
	array := make([]interface{}, len(dataSlice))
	for i, v := range dataSlice {
		array[i] = v
	}
	// make heap
	heap := Heap{array, LessThan}
	heap.BuildHeap()
	// test Size()
	if heap.Size() != 5 {
		t.Fatal("Size() does not return correct size")
	}
	// test Clear()
	heap.Clear()
	if heap.Size() != 0 {
		fmt.Printf("size is %d\n", heap.Size())
		t.Fatal("clear() does not work properly")
	}
	// test again after Clear() being called
	heap.Clear()
	heap.data = array
	if heap.Size() != 5 {
		t.Fatal("Size() does not return correct size")
	}
}

// Return true if slice A is sorted according to comparison function f
func IsSorted(A []interface{}, f container.CompFunc) bool {
	for i := 0; i < len(A)-1; i++ {
		if f(A[i+1], A[i]) != true {
			return false
		}
	}

	return true
}

// Test HeapSort() function in heap.go
func TestHeapSort(t *testing.T) {
	fmt.Println("TestBuildHeap")
	dataSlice := []int{27, 17, 3, 16, 13, 10, 1, 5, 7, 12, 4, 8, 9, -1}
	array := make([]interface{}, len(dataSlice))
	for i, v := range dataSlice {
		array[i] = v
	}
	// make max heap
	heap := Heap{array, MoreThan}
	heap.HeapSort()
	if IsSorted(heap.data, heap.comp) == false {
		fmt.Println(heap.data)
		t.Fatal("The data array is not sorted as expected")
	}

	heap = Heap{array, LessThan}
	heap.HeapSort()
	if IsSorted(heap.data, heap.comp) == false {
		fmt.Println(heap.data)
		t.Fatal("The data array is not sorted as expected")
	}

}
