package main

import (
	"fmt"
	. "main/linkedlist"
	"math/rand"
	"time"
)

func main() {
	var ll LinkedList[int]

	seed := time.Now().UnixNano()
	random := rand.New(rand.NewSource(seed))

	// Insert Nodes
	NODE_COUNT := 10
	min := 1
	max := 10
	for i := 0; i < NODE_COUNT; i++ {
		r := random.Intn(max-min+1) + min
		ll.Insert(&Node[int]{Data: r})
	}

	fmt.Println(ll.String())

	ll1, ll2 := ll.SplitBy(ll.Mid())

  fmt.Printf("\nSplit 1\nHead: %s\nSize: %d\nMid: %s\n", ll1.Head(), ll1.Size(), ll1.Mid())
	fmt.Println(ll1.String())

  fmt.Printf("\nSplit 2\nHead: %s\nSize: %d\nMid: %s\n", ll2.Head(), ll2.Size(), ll2.Mid())
  fmt.Println(ll2.String())
}
