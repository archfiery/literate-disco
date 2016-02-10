package container

import (
	"fmt"
	"testing"
)

type ContainerTestStruct struct {
	data []interface{}
}

func (c *ContainerTestStruct) Clear() {
	c.data = make([]interface{}, 0)
}

func (c ContainerTestStruct) Empty() bool {
	if c.Size() == 0 {
		return true
	}
	return false
}

func (c ContainerTestStruct) Size() int {
	return len(c.data)
}

func (c *ContainerTestStruct) Insert(i int) {
	c.data = append(c.data, i)
}

func TestContainerClear(t *testing.T) {
	fmt.Println("TestContainerClear")
	c := ContainerTestStruct{}
	c.Insert(1)
	c.Insert(1)
	c.Insert(1)
	if c.Size() != 3 {
		t.Fatal("Size does not match, expected 3, but get ", c.Size())
	}
	c.Clear()
	if c.Size() != 0 {
		t.Fatal("Size does not match, expected 0, but get ", c.Size())
	}
}

func TestContainerEmpty(t *testing.T) {
	fmt.Println("TestContainerEmpty")
	c := ContainerTestStruct{}
	c.Insert(1)
	if c.Empty() != false {
		t.Fatal("c should not be empty")
	}
	c.Clear()
	if c.Empty() != true {
		t.Fatal("c should be empty, but it has size ", c.Size())
	}
}

func TestContainerSize(t *testing.T) {
	fmt.Println("TestContainerSize")
	c := ContainerTestStruct{}
	if c.Size() != 0 {
		t.Fatal("Size does not match, expected 0, but get ", c.Size())
	}
	c.Insert(1)
	c.Insert(1)

	if c.Size() != 2 {
		t.Fatal("Size does not match, expected 2, but get ", c.Size())
	}
}
