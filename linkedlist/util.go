package linkedlist

// Mid Finds midpoint node in a linked-list
func Mid[T any](head *Node[T], tails ...*Node[T]) *Node[T] {
	var tail *Node[T]
	if len(tails) > 0 {
		tail = tails[0]
	}

	slow, fast := head, head
	for ; fast != tail && fast.Next != tail; slow, fast = slow.Next, fast.Next.Next {
	}
	return slow
}

func IsEqual[T comparable](ll1, ll2 *LinkedList[T]) bool {
	if ll1.Size() != ll2.Size() {
		return false
	}

	node1 := ll1.Head
	node2 := ll2.Head

	var prev *Node[T]
	for ; prev != ll1.Tail; prev, node1, node2 = node1, node1.Next, node2.Next {
		if node1.Data != node2.Data {
			return false
		}
	}

	return true
}
