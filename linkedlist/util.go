package linkedlist

// Mid finds midpoint in a linkedlist
func Mid[T any](head *Node[T]) *Node[T] {
	slow, fast := head, head
	for ; fast != nil && fast.next != nil; slow, fast = slow.next, fast.next.next {
	}
	return slow
}

func IsEqual[T any](ll1, ll2 LinkedList[T]) bool {
  panic("Not Implemented")
}