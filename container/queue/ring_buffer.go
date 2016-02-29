package queue

import (
	"github.com/archfiery/literate-disco/error"
	"runtime"
	"sync/atomic"
)

const (
	cacheline_size   = 64
	freeup_threshold = 10000
)

// The padding array
type cacheline_pad [cacheline_size]byte

// The ring_cell struct represents the data cell in the ring buffer
type ring_cell struct {
	sequence_ uint64
	data_     interface{}
}

// The array of ring_cell pointers
type buffer []*ring_cell

// Type RingBuffer represents a ring buffer data structure.
// It has 2 primary operations, `enqueue` and `dequeue`.
// Each `enqueue` and `dequeue` takes 1 CAS per operation.
// The original implementation is from http://www.1024cores.net/home/lock-free-algorithms/queues/bounded-mpmc-queue
// Some additional ones are learnt from https://github.com/Workiva/go-datastructures/blob/master/queue/ring.go
type RingBuffer struct {
	pad0_        cacheline_pad
	buffer_      buffer
	buffer_mask_ uint64
	pad1_        cacheline_pad
	enqueue_pos_ uint64
	pad2_        cacheline_pad
	dequeue_pos_ uint64
	pad3_        cacheline_pad
}

// Function roundUp returns the nearest power of 2
func roundUp(v uint64) uint64 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v |= v >> 32
	v++
	return v
}

// MakeRingBuffer returns an initialized RingBuffer
// Supplied size will be rounded up the nearest power of 2
// If the size is not satisfied, an AssertionError will be returned
func MakeRingBuffer(buffer_size uint64) (*RingBuffer, error.Error) {
	buffer_size = roundUp(buffer_size)
	rb := RingBuffer{}
	// assert the buffer_size is 2 ^ x
	if buffer_size >= 2 && buffer_size&(buffer_size-1) == 0 {
	} else {
		return &rb, error.AssertionError{}
	}

	rb.buffer_ = make(buffer, buffer_size)
	var i uint64
	for i = 0; i < buffer_size; i++ {
		rb.buffer_[i] = &ring_cell{sequence_: i}
	}
	rb.buffer_mask_ = buffer_size - 1
	atomic.StoreUint64(&rb.enqueue_pos_, 0)
	atomic.StoreUint64(&rb.dequeue_pos_, 0)

	return &rb, nil
}

// Enqueue adds a new element to the tail of the ring buffer
// It returns true if the operation is successful, false otherwise
// It blocks on a full queue
func (rb *RingBuffer) Enqueue(data interface{}) bool {
	var cell *ring_cell
	pos := atomic.LoadUint64(&rb.enqueue_pos_)
	i := 0
Loop:
	for {
		cell = rb.buffer_[pos&rb.buffer_mask_]
		seq := atomic.LoadUint64(&cell.sequence_)
		switch dif := seq - pos; {
		case dif == 0:
			if atomic.CompareAndSwapUint64(&rb.enqueue_pos_, pos, pos+1) {
				break Loop
			}
		case dif < 0:
			return false
		default:
			pos = atomic.LoadUint64(&rb.enqueue_pos_)
		}
		// freeup the cpu
		if i >= freeup_threshold {
			runtime.Gosched()
			i = 0
		} else {
			i++
		}
	}

	cell.data_ = data
	atomic.StoreUint64(&cell.sequence_, pos+1)
	return true
}

// Dequeue removes and returns the `oldest` element from the ring buffer
// It also returns true if the operation is successful, false otherwise
// It blocks on an empty queue
func (rb *RingBuffer) Dequeue() (data interface{}, b bool) {
	var cell *ring_cell
	pos := atomic.LoadUint64(&rb.dequeue_pos_)
	i := 0
Loop:
	for {
		cell = rb.buffer_[pos&rb.buffer_mask_]
		seq := atomic.LoadUint64(&cell.sequence_)
		switch dif := seq - pos - 1; {
		case dif == 0:
			if atomic.CompareAndSwapUint64(&rb.dequeue_pos_, pos, pos+1) {
				break Loop
			}
		case dif < 0:
			return nil, false
		default:
			pos = atomic.LoadUint64(&rb.dequeue_pos_)
		}
		// freeup the cpu
		if i >= freeup_threshold {
			runtime.Gosched()
			i = 0
		} else {
			i++
		}
	}
	data = cell.data_
	atomic.StoreUint64(&cell.sequence_, pos+rb.buffer_mask_+1)
	b = true
	return data, b
}

// Returns the current size of the ring buffer
func (rb RingBuffer) Size() uint64 {
	return atomic.LoadUint64(&rb.enqueue_pos_) - atomic.LoadUint64(&rb.dequeue_pos_)
}

// Synonym of method Size()
func (rb RingBuffer) Len() uint64 {
	return rb.Size()
}

// Returns the capacity of the ring buffer
func (rb RingBuffer) Cap() uint64 {
	return uint64(len(rb.buffer_))
}
