package linkedlist

import (
	"fmt"
	"strings"
)

type LinkedList[T any] struct {
  _ struct{}
	head *Node[T]
	tail *Node[T]
	mid  *Node[T]
	size uint
}

// Return a copy of head so that the internal linked list
// cannot be manipulated directly
func (ll *LinkedList[T]) Head() *Node[T] {
	headCopy := Node[T]{
		Data: ll.head.Data,
		next: ll.head.next,
	}
	return &headCopy
}

func (ll *LinkedList[T]) Insert(n *Node[T]) *LinkedList[T] {
	if ll.head == nil {
		ll.head = n
		ll.tail = ll.head
		ll.mid = ll.head
	} else {
		ll.tail.next = n
		ll.tail = n
	}
	ll.size++

	// Odd
	if ll.size%2 != 0 && n != ll.head {
		ll.mid = ll.mid.next
	}

	fmt.Println("Size: ", ll.size)
	fmt.Println("Mid: ", ll.mid.String())
	fmt.Println()

	return ll
}

func (ll *LinkedList[T]) Size() uint {
	return ll.size
}

func (ll *LinkedList[T]) String() string {
	var str strings.Builder

	// Iterate through the list one by one
  ll.MapFilter(func (ptr *Node[T]) *Node[T] {
    if ptr == ll.head {
			str.WriteString(fmt.Sprintf("[HEAD] → "))
		} else {
			str.WriteString(fmt.Sprintf(" → "))
		}

		str.WriteString(ptr.String())

		if ptr == ll.tail {
			str.WriteString(fmt.Sprintf(" → [NULL]"))
		}
    return ptr
  })

	return str.String()
}

func (ll *LinkedList[T]) Mid() *Node[T] {
	if ll.mid != nil {
		return ll.mid
	}

	ll.mid = Mid(ll.head)

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

	head1 = ll.head
	tail1 = node

	head2 = node.next

  if keepMid {
    head2 = node
  }
	tail2 = ll.tail

	size1 = ll.size / 2
	if ll.size%2 == 0 || keepMid {
		size2 = ll.size / 2
	} else {
		size2 = (ll.size + 1) / 2
	}

	ll1 := LinkedList[T]{
		head: head1,
		tail: tail1,
		size: size1,
	}

	ll2 := LinkedList[T]{
		head: head2,
		tail: tail2,
		size: size2,
	}

	return &ll1, &ll2
}

// Reverse linked list in place
func (ll *LinkedList[T]) Reverse() *LinkedList[T] {
	panic("Not Implemented")
}

func (ll *LinkedList[T]) IsPalindrome() bool {
	panic("Not Implemented")
}

// MapFilter will map the result of the func passed.
// If return is `nil`, then it will drop the node.
// Can be used to iterate a LinkedList
func (ll *LinkedList[T]) MapFilter(fn func (n *Node[T]) *Node[T]) *LinkedList[T] {
  var prev *Node[T]
  for ptr := ll.head; prev != ll.tail; prev, ptr = ptr, ptr.next {
    res := fn(ptr)
    if res != nil {
      ptr = res
    } else if prev != nil /* && res == nil */ {
      // Drop the curr node (ptr)
      prev.next = ptr.next

      // Remove any further connections from curr node (ptr)
      ptr.next = nil
    }
  }
  return ll
}