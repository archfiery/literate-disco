package list

import "github.com/archfiery/literate-disco/error"

// A Node struct that represents a node in the list
type Node struct {
	prev *Node
	item interface{}
	next *Node
}

// Returns a node with default value
func MakeNode(next *Node, e interface{}, prev *Node) Node {
	n := Node{next, e, prev}
	return n
}

// A LinkedList struct that represents a linked list
// It maintains the first and last node of the list
// It also maintains the size of the list
// The actual links are stored in the node
type LinkedList struct {
	size int
	first *Node
	last  *Node
}

// Returns a newly made list
// first and last nodes are nil by default
func MakeList() LinkedList {
	return LinkedList{0, nil, nil}
}

//==========
// Capacity
//==========

// Returns true if the linked list is empty
func (list LinkedList) Empty() bool {
	return list.size == 0
}

// Returns the size of the linked list
func (list LinkedList) Size() int {
	return list.size
}

//================
// Element Access
//================

// Returns the first element in the list
// Returns NoSuchElementError if the list is empty
func (list LinkedList) Front() (interface{}, error.Error) {
	f := list.first
	if f == nil {
		return -1, error.NoSuchElementError{}
	}
	return f.item, nil
}

// Returns the last element in the list
// Returns NoSuchElementError if the list is empty
func (list LinkedList) Back() (interface{}, error.Error) {
	l := list.last
	if l == nil {
		return -1, error.NoSuchElementError{}
	}
	return l.item, nil
}

//===========
// Modifiers
//===========

// Removes the first node from the list and returns its element
// Returns NoSuchElementError if the list is empty
func (list *LinkedList) PopFront() (interface{}, error.Error) {
	f := list.first
	if f == nil {
		return -1, error.NoSuchElementError{}
	}
	return list.unlinkFirst(f), nil
}

// Removes the last node from the list and returns its element
// Returns NoSuchElementError if the list is empty
func (list *LinkedList) PopBack() (interface{}, error.Error) {
	l := list.last
	if l == nil {
		return -1, error.NoSuchElementError{}
	}
	return list.unlinkLast(l), nil
}

// Inserts an element at the beginning of the list
func (list *LinkedList) PushFront(e interface{}) {
	list.linkFirst(e)
}

// Inserts an element at the end of the list
func (list *LinkedList) PushBack(e interface{}) {
	list.linkLast(e)
}
