package heap

import (
	"fmt"
	"testing"
)

func comp(a interface{}, b interface{}) bool {
	return a.(int) < b.(int)
}

func TestHeapify(t *testing.T) {
	fmt.Println("TestHeapify")
	// make dummy data
	dataSlice := []int{3, 1, 2}
	array := make([]interface{}, len(dataSlice))
	for i, v := range dataSlice {
		array[i] = v
	}
	// make heap
	heap := Heap{array, comp}
	// verify heap size
	if heap.Size() != 3 {
		t.Fatal("The size does not match")
	}

	heap.Heapify(0)
	if heap.At(0) != 1 {
		t.Fatal("Heapify error")
	}
}

func verifyHeapProperty(heap Heap) bool {
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

func TestBuildHeap(t *testing.T) {
	fmt.Println("TestBuildHeap")
	dataSlice := []int{27, 17, 3, 16, 13, 10, 1, 5, 7, 12, 4, 8, 9, 0}
	array := make([]interface{}, len(dataSlice))
	for i, v := range dataSlice {
		array[i] = v
	}
	// make heap
	heap := Heap{array, comp}
	heap.BuildHeap()

	// expected answer for test array
	expected := []int{0, 4, 1, 5, 12, 8, 3, 16, 7, 17, 13, 10, 9, 27}
	for i, v := range heap.data {
		if v != expected[i] {
			fmt.Println("Answer is not expected")
		}
	}

	if verifyHeapProperty(heap) != true {
		fmt.Println("Heap property violated")
	}
}
