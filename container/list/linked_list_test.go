package list

import (
	"fmt"
	"github.com/archfiery/literate-disco/test"
	"testing"
)

// A simple test that tests node creation
func TestCreateNode(t *testing.T) {
	fmt.Println("\nTestCreateNode LinkedList")
	n1 := MakeNode(nil, 1, nil)
	if n1.next != nil || n1.prev != nil {
		t.Fatal("Links on the node is incorrect")
	}
}

// A simple test that contains `weak test cases` for the linked list
func TestSimpleLinkedList(t *testing.T) {
	fmt.Println("TestCreateLinkedList LinkedList")
	l := MakeList(test.IntEqual)
	// push 1 to back, 1
	l.PushBack(1)
	if val, err := l.Front(); err == nil && val != 1 {
		t.Fatal("The first node is expected to return 1, but it returns ", val)
	}
	if val, err := l.Back(); err == nil && val != 1 {
		t.Fatal("The last node is expected to return 1, but it returns ", val)
	}
	// push 2 to back 1 <-> 2
	l.PushBack(2)
	if val, err := l.Front(); err == nil && val != 1 {
		t.Fatal("The first node is expected to return 1, but it returns ", val)
	}
	if val, err := l.Back(); err == nil && val != 2 {
		t.Fatal("The last node is expected to return 2, but it returns ", val)
	}
	// push 3 to back 1 <-> 2 <-> 3
	l.PushBack(3)
	// push 4 to back 1 <-> 2 <-> 3 <-> 4
	l.PushBack(4)
	// pop front 2 <-> 3 <-> 4
	l.PopFront()
	if l.Size() != 3 {
		t.Fatal("The size is expected to be 3, but it is ", l.Size())
	}
	if val, err := l.Front(); err == nil && val != 2 {
		t.Fatal("The first node is expected to return 2, but it returns ", val)
	}
	if val, err := l.Back(); err == nil && val != 4 {
		t.Fatal("The last node is expected to return 4, but it returns ", val)
	}
}

// Test the indexOf helper method
func TestIndexOf(t *testing.T) {
	fmt.Println("TestIndexOf")
	l := MakeList(test.IntEqual)
	// push 1 to back 1
	l.PushBack(1)
	// push 2 to back 1 <-> 2
	l.PushBack(2)
	if val := l.indexOf(2); val != 1 {
		t.Fatal("Index of 2 is expected to be 1, but it is ", val)
	}
	// push 2 to back 1 <-> 2 <-> 2
	l.PushBack(2)
	if val := l.indexOf(2); val != 1 {
		t.Fatal("Index of 2 is expected to be 1, but it is ", val)
	}
	// push 2 to back 1 <-> 2 <-> 2 <-> 2
	l.PushFront(2)
	if val := l.indexOf(2); val != 0 {
		t.Fatal("Index of 2 is expected to be 0, but it is ", val)
	}
	// push 15 to back 1 <-> 2 <-> 2 <-> 2 <-> 15
	l.PushBack(15)
	if val := l.indexOf(15); val != 4 {
		t.Fatal("Index of 15 is expected to be 4, but it is ", val)
	}
}

// Test the lastIndexOf helper method
func TestLastIndexOf(t *testing.T) {
	fmt.Println("TestLastIndexOf")
	l := MakeList(test.IntEqual)
	// push 1 to back 1
	l.PushBack(1)
	// push 2 to back 1 <-> 2
	l.PushBack(2)
	if val := l.lastIndexOf(2); val != 1 {
		t.Fatal("Last index of 2 is expected to be 1, but it is ", val)
	}
	// push 2 to back 1 <-> 2 <-> 2
	l.PushBack(2)
	if val := l.lastIndexOf(2); val != 2 {
		t.Fatal("Last index of 2 is expected to be 1, but it is ", val)
	}
	// push 15 to back and front 15 <-> 1 <-> 2 <-> 2 <-> 15
	l.PushBack(15)
	l.PushFront(15)
	if val := l.indexOf(15); val != 0 {
		t.Fatal("Index of 15 is expected to be 4, but it is ", val)
	}
	if val := l.lastIndexOf(15); val != 4 {
		t.Fatal("Last index of 15 is expected to be 4, but it is ", val)
	}
}

// Test Insert method
func TestInsert(t *testing.T) {
	fmt.Println("TestInsert")
	l := MakeList(test.IntEqual)
	if err := l.Insert(1, 15); err == nil {
		t.Fatal("An OutOfRangeError should be occured")
	}
	if err := l.Insert(1, 0); err == nil {
		if val, _ := l.Front(); val != 1 {
			t.Fatal("Insert is not successful at the front of the list")
		}
		if val, _ := l.Back(); val != 1 {
			t.Fatal("Insert is not successful at the front of the list")
		}
	} else {
		t.Fatal(err.Error())
	}
	if err := l.Insert(2, 1); err == nil {
		if val, _ := l.Front(); val != 1 {
			t.Fatal("Insert is not successful at the end of the list")
		}
		if val, _ := l.Back(); val != 2 {
			t.Fatal("Insert is not successful at the end of the list")
		}
	} else {
		t.Fatal(err.Error())
	}
	l.PopBack()
	l.PopBack()
	if l.Size() != 0 {
		t.Fatal("Size is expected to be 0, but it is ", l.Size())
	}
	// bulk push
	for i := 0; i < 10; i++ {
		l.PushBack(i)
	}
	l.Insert(-22, 5)
	if val := l.indexOf(-22); val != 5 {
		t.Fatal("index of -22 is expected to be 5, but it is ", val)
	}
	l.Insert(-23, 5)
	if val := l.indexOf(-22); val != 6 {
		t.Fatal("index of -22 is expected to be 6, but it is ", val)
	}
}

// Test Erase method
func TestErase(t *testing.T) {
	fmt.Println("TestErase")
	l := MakeList(test.IntEqual)
	err := l.Erase(0)
	if err == nil {
		fmt.Println("OutOfRangeError is expected to occur")
		t.Fatal(err.Error())
	}
	// push 1 to back 1
	l.PushBack(1)
	// push 2 to back 1 <-> 2
	l.PushBack(2)
	// push 3 to the back 1 <-> 2 <-> 3
	l.PushBack(3)
	// pop at pos 0
	err = l.Erase(0)
	if err != nil {
		t.Fatal(err.Error())
	}
	// 2 <-> 3
	if val, _ := l.Front(); val != 2 {
		t.Fatal("Front value is expected to be 2, but it is ", val)
	}
	if val, _ := l.Back(); val != 3 {
		t.Fatal("Back value is expected to be 3, but it is ", val)
	}
	// pop at pos 0
	err = l.Erase(0)
	if err != nil {
		t.Fatal(err.Error())
	}
	// 3
	if val, _ := l.Front(); val != 3 {
		t.Fatal("Front value is expected to be 3, but it is ", val)
	}
	if val, _ := l.Back(); val != 3 {
		t.Fatal("Back value is expected to be 3, but it is ", val)
	}
	// then push 4 5 6, expected: 6 <-> 5 <-> 4 <-> 3
	l.PushFront(4)
	l.PushFront(5)
	l.PushFront(6)
	// pop at pos 2, expected: 6 <-> 5 <-> 3
	l.Erase(2)
	if l.Size() != 3 {
		t.Fatal("Size of list is expected to be 3, but it is ", l.Size())
	}
	if val := l.indexOf(5); val != 1 {
		t.Fatal("index of 5 is expected to be 1, but it ", val)
	}
	// Erase at pos 0 until the list is empty
	for l.Size() > 0 {
		l.Erase(0)
	}
	if l.Size() != 0 {
		t.Fatal("Size of list is expected to be 0, but it is ", l.Size())
	}

}
