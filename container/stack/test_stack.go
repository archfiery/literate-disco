package stack

import (
	"fmt"
	"testing"
	"github.com/archfiery/literate-disco/test"
)

func TestCreateStack(t *testing.T) {
	fmt.Println("TestCreateStack")
	s := MakeStack(sort.LessThan)
	if s.Size() != 0 {
		t.Fatal("The size of stack is not 0")
	}
}

func TestPush(t *testing.T) {
	fmt.Println("TestPush")
	s := MakeStack(sort.LessThan)
	s.Push(1)
	s.Push(2)
	if s.Size() != 2 {
		t.Fatal("Size is incorrect after 2 pushes")
	}
}

func TestPop(t *testing.T) {
	fmt.Println("TestPop")
	s := MakeStack(sort.LessThan)
	s.Push(1)
	s.Pop()
	if s.Size() != 0 {
		t.Fatal("Size is incorrect after 1 push and 1 pop")
	}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Pop()
	if s.Size() != 2 {
		t.Fatal("Size is incorrect after 3 push and 1 pop")
	}

}

func TestTop(t *testing.T) {
	fmt.Println("TestTop")
	s := MakeStack(sort.LessThan)
	s.Push(1)
	s.Push(2)
	val, err := s.Top()
	if err != nil || val != 2 {
		t.Fatal("Top does not return correct answer, expected 2, but get", val)
	}
	s.Pop()
	val, err = s.Top()
	if err != nil || val != 1 {
		t.Fatal("Top does not return correct answer, expected 1, but get", val)
	}
	s.Pop()

	val, err = s.Top()
	if err == nil {
		t.Fatal("Running top on an empty stack should return error")
	}
}
