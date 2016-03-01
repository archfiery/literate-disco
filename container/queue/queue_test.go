package queue

import (
	"fmt"
	"testing"
)

func TestMakeQueue(t *testing.T) {
	fmt.Println("TestMakeQueue")
	q := MakeQueue()
	if q.Capacity() != INIT_CAP {
		t.Fatal("Capacity of the queue is incorrect")
	}
	if q.Size() != 0 {
		t.Fatal("Size of the queue is incorrect")
	}
}

func TestPopFromEmptyQueue(t *testing.T) {
	fmt.Println("TestPopFromEmptyQueue")
	q := MakeQueue()
	err := q.Pop()
	if err == nil {
		t.Fatal("Expect to have error returned, but nil")
	}
}

func TestPushToQueue(t *testing.T) {
	fmt.Println("TestPushToQueue")
	q := MakeQueue()
	for i := 0; i < 3000; i++ {
		q.Push(i)
	}
	if q.Capacity() != INIT_CAP*8 {
		t.Fatal("Capacity is expected to be ", INIT_CAP*8, ", but it is ", q.Capacity())
	}
	if q.Size() != 3000 {
		t.Fatal("Size is expected to be 3000, but it is ", q.Size())
	}
}

func TestPopFromQueue(t *testing.T) {
	fmt.Println("TestPopFromQueue")
	q := MakeQueue()
	for i := 0; i < 10; i++ {
		q.Push(i)
	}
	for i := 0; i < 3; i++ {
		q.Pop()
	}
	if val, _ := q.Front(); val != 3 {
		t.Fatal("Error on Front")
	}
	if val, _ := q.Back(); val != 9 {
		t.Fatal("Error on Back")
	}
}
