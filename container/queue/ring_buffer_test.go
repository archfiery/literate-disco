package queue

import (
	"fmt"
	"sync"
	"testing"
)

func TestRingBufferEnqueue(t *testing.T) {
	fmt.Println("\nTestRingBufferEnqueue")
	rb, err := MakeRingBuffer(6)
	if err != nil || len(rb.buffer_) != 8 {
		t.Fatal("The ring buffer size is not correctly set")
	}

	succ := rb.Enqueue(1)
	if succ == false {
		t.Fatal("Fail to enqueue")
	}

	if rb.Size() != 1 {
		t.Fatal("Size is incorrect")
	}

	val, succ := rb.Dequeue()
	if succ == false {
		t.Fatal("Fail to dequeue")
	}
	if val != 1 {
		t.Fatal("Value is incorrect")
	}

	if rb.Size() != 0 {
		t.Fatal("Size is incorrect")
	}
}

func TestRingBufferEnqueueDequeue(t *testing.T) {
	fmt.Println("TestRingBufferEnqueueDequeue")
	rb, err := MakeRingBuffer(9)
	if err != nil || len(rb.buffer_) != 16 {
		t.Fatal("The ring buffer size is not correctly set")
	}

	rb.Enqueue(1)
	rb.Enqueue(2)
	val, _ := rb.Dequeue()
	if val != 1 {
		t.Fatal("Incorrect dequeued value, expected 1, but it is ", val)
	}
	rb.Enqueue(15)
	val, _ = rb.Dequeue()
	if val != 2 {
		t.Fatal("Incorrect dequeued value, expected 2, but it is ", val)
	}
	val, _ = rb.Dequeue()
	if val != 15 {
		t.Fatal("Incorrect dequeued value, expected 15, but it is ", val)
	}
	rb.Enqueue(12)
	if rb.Size() != 1 || rb.Cap() != 16 {
		t.Fatal("Expected size is 1, but it is  ", rb.Size(), ". Expected cap is ", rb.Cap())
	}
}

//==============
// Benchmark
//==============

func BenchmarkSinglePCRingBuffer(b *testing.B) {
	rb, _ := MakeRingBuffer(1024)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			rb.Enqueue(i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			rb.Dequeue()
		}
	}()

	wg.Wait()
}

func BenchmarkBenchGet(b *testing.B) {
	rbs := make([]*RingBuffer, 0, b.N)

	for i := 0; i < b.N; i++ {
		rb, _ := MakeRingBuffer(2)
		rbs = append(rbs, &rb)
	}

	b.ResetTimer()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			rbs[i].Enqueue(1)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			rbs[i].Dequeue()
		}
	}()

	wg.Wait()
}
