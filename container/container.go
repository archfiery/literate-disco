package container

type Container interface {
	Clear()
	Empty() bool
	Size() int
}
