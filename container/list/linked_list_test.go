package list

import (
	"testing"
	"fmt"
)

func TestCreateNode(t *testing.T) {
	fmt.Println("TestCreateNode LinkedList")
	n1 := MakeNode(nil, 1, nil)
	if n1.next != nil || n1.prev != nil {
		t.Fatal("Links on the node is incorrect")
	}
}

// A simple test that contains `weak test cases` for the linked list
func TestSimpleLinkedList(t *testing.T) {
	fmt.Println("TestCreateLinkedList LinkedList")
	l := MakeList()
	l.PushBack(1)
	val, err := l.Front()
	if err == nil && val != 1 {
		t.Fatal("The first node is expected to return 1, but it returns ", val)
	}
	val, err = l.Back()
	if err == nil && val != 1 {
		t.Fatal("The last node is expected to return 1, but it returns ", val)
	}
	l.PushBack(2)
	if err == nil && val != 2 {
		t.Fatal("The last node is expected to return 2, but it returns ", val)
	}
	l.PushBack(3)
	l.PushBack(4)
	l.PopFront()
	if l.Size() != 3 {
		t.Fatal("The size is expected to be 3, but it is ", l.Size())
	}
	val, err = l.Front()
	if err == nil && val != 2 {
		t.Fatal("The first node is expected to return 2, but it returns ", val)
	}
}
