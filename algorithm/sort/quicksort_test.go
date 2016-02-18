package sort

import (
	"fmt"
	"testing"
)

func TestQuicksortPartition(t *testing.T) {
	fmt.Println("TestQuicksortPartition")
	dataSlice := []int{27, 17, 3, 16, 13}
	array := make([]interface{}, len(dataSlice))
	for i, v := range dataSlice {
		array[i] = v
	}

	val, err := quicksortPartition(array, 0, 4, leq)
	if err != nil {
		t.Fatal("quick sort error")
	}
	if val != 1 {
		t.Fatal("pivot position is incorrect")
	}

	fmt.Println(array)
}
