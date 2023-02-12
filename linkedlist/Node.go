package linkedlist

import "fmt"

type Node[T any] struct {
	_    struct{}
	Data T
	Next *Node[T]
}

func (n *Node[T]) String() string {
	return fmt.Sprintf("%v", n.Data)
}
