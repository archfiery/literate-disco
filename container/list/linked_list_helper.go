package list


//===============
// helper methods
//===============

// Creates a node and makes it the first element in the list
func (list *LinkedList) linkFirst(e interface{}) {
	f := list.first
	n := MakeNode(nil, e, f)
	if f == nil {
		list.last = &n
	} else {
		f.prev = &n
	}
	list.size++
}

// Creates a node and makes it the last element in the list
func (list *LinkedList) linkLast(e interface{}) {
	l := list.last
	n := MakeNode(l, e, nil)
	if l == nil {
		list.first = &n
	} else {
		l.next = &n
	}
	list.size++
}

// Creates a node and insert it before some node in the list
func (list *LinkedList) linkBefore(e interface{}, succ *Node) {
	pred := succ.prev
	n := MakeNode(pred, e, nil)
	succ.prev = &n
	if (pred == nil) {
		list.first = &n
	} else {
		pred.next = &n
	}
	list.size++
}

// Unlink the first node that is not nil
func (list *LinkedList) unlinkFirst(f *Node) interface{} {
	// It does not check if the f is the first node or not
	// assert list.first = f && f != nil
	// cache the item the first node holds
	e := f.item
	next := f.next
	// clear the first node
	f.item = nil
	f.next = nil

	list.first = next
	if next == nil {
		list.last = nil
	} else {
		next.prev = nil
	}
	list.size--

	return e
}

// Unlink the last node that is not nil
func (list *LinkedList) unlinkLast(l *Node) interface{} {
	// It does not check if the l is the last node or not
	// assert list.last = l && l != nil
	e := l.item
	prev := l.prev
	// clear the last node
	l.item = nil
	l.prev = nil

	list.last = prev
	if prev == nil {
		list.first = nil
	} else {
		prev.next = nil
	}
	list.size--

	return e
}

// Unlinks a node that is not nil
func (list *LinkedList) unlink(n *Node) interface{} {
	// It does not check if the n is nil or not
	// assert n != nil
	e := n.item
	prev := n.prev
	next := n.next

	if prev == nil {
		list.first = next
	} else {
		prev.next = next
		n.prev = nil
	}

	if next == nil {
		list.last = prev
	} else {
		next.prev = prev
		n.next = nil
	}

	n.item = nil
	list.size--

	return e
}

