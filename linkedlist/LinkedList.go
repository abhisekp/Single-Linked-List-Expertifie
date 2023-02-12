package linkedlist

import (
	"fmt"
	"strings"
)

type LinkedList[T comparable] struct {
	_    struct{}
	Head *Node[T]
	Tail *Node[T]
	mid  *Node[T]
	size uint
}

func (ll *LinkedList[T]) updateSize() {
	var prev *Node[T]
	for ptr := ll.Head; prev != ll.Tail; prev, ptr = ptr, ptr.Next {
		ll.size++
	}
}

func (ll *LinkedList[T]) Size() uint {
	if ll.size == 0 && ll.Head != nil {
		ll.updateSize()
	}
	return ll.size
}

func (ll *LinkedList[T]) Insert(n *Node[T]) *LinkedList[T] {
	if ll.Head == nil {
		ll.Head = n
		ll.Tail = ll.Head
		ll.mid = ll.Head
	} else {
		ll.Tail.Next = n
		ll.Tail = n
	}
	ll.size++

	// Odd
	if ll.size%2 != 0 && n != ll.Head {
		ll.mid = ll.mid.Next
	}

	// fmt.Println("Size: ", ll.Size)
	// fmt.Println("Mid: ", ll.mid.String())
	// fmt.Println()

	return ll
}

func (ll *LinkedList[T]) String() string {
	var str strings.Builder

	// Iterate through the list one by one
	ll.MapFilter(func(ptr *Node[T]) *Node[T] {
		if ptr == ll.Head {
			str.WriteString("[HEAD] → ")
		} else {
			str.WriteString(" → ")
		}

		str.WriteString(ptr.String())

		if ptr == ll.Tail {
			str.WriteString(" → [NULL]")
		}
		return ptr
	})

	return str.String()
}

func (ll *LinkedList[T]) Details() string {
	var detail strings.Builder
	detail.WriteString(fmt.Sprintf("Head: %s\nTail: %s\nSize: %d\nMid: %s", ll.Head, ll.Tail, ll.size, ll.mid))
	return detail.String()
}

func (ll *LinkedList[T]) updateMid() {
	ll.mid = Mid(ll.Head, ll.Tail)
}

func (ll *LinkedList[T]) Mid() *Node[T] {
	if ll.mid != nil {
		return ll.mid
	}

	ll.updateMid()

	return ll.mid
}

func (ll *LinkedList[T]) SplitBy(node *Node[T], keepMids ...bool) (*LinkedList[T], *LinkedList[T]) {
	var head1, head2 *Node[T]
	var tail1, tail2 *Node[T]
	var size1, size2 uint

	keepMid := false
	if len(keepMids) > 0 {
		keepMid = keepMids[0]
	}

	head1 = ll.Head
	tail1 = node

	head2 = node.Next

	if keepMid {
		head2 = node
	}
	tail2 = ll.Tail

	size1 = ll.size / 2
	if ll.size%2 == 0 || keepMid {
		size2 = ll.size / 2
	} else {
		size2 = (ll.size + 1) / 2
	}

	ll1 := LinkedList[T]{
		Head: head1,
		Tail: tail1,
		size: size1,
	}
	ll1.Mid()

	ll2 := LinkedList[T]{
		Head: head2,
		Tail: tail2,
		size: size2,
	}
	ll2.Mid()

	return &ll1, &ll2
}

// Reverse linked list in place
func (ll *LinkedList[T]) Reverse() *LinkedList[T] {
	var prev, ptr *Node[T]

	for ptr = ll.Head; ptr != ll.Tail; {
		next := ptr.Next
		ptr.Next = prev
		prev = ptr
		ptr = next
	}

	ll.Tail = ll.Head
	ll.Head = ptr

	return ll
}

func (ll *LinkedList[T]) IsPalindrome() bool {
	ll1, ll2 := ll.SplitBy(ll.Mid(), true)
	ll2.Reverse()
	isEqual := IsEqual(ll1, ll2)
	ll2.Reverse()
	ll1.Tail.Next = ll2.Head
	return isEqual
}

// MapFilter will map the result of the func passed.
// If return is `nil`, then it will drop the node.
// Can be used to iterate a LinkedList
func (ll *LinkedList[T]) MapFilter(fn func(n *Node[T]) *Node[T]) *LinkedList[T] {
	var prev *Node[T]
	for ptr := ll.Head; prev != ll.Tail; prev, ptr = ptr, ptr.Next {
		res := fn(ptr)
		if res != nil {
			ptr = res
		} else if prev != nil /* && res == nil */ {
			// Drop the curr node (ptr)
			prev.Next = ptr.Next

			// Remove any further connections from curr node (ptr)
			ptr.Next = nil
		}
	}
	return ll
}

func (ll *LinkedList[T]) Merge(ll2 *LinkedList[T]) *LinkedList[T] {
	if ll2 == nil {
		return ll
	}

	// Update Head of ll if nil
	if ll.Head == nil {
		ll.Head = ll2.Head
	}

	// Update Tail of ll
	ll.Tail.Next = ll2.Head

	ll.updateMid()
	ll.updateSize()

	return ll
}
