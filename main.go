package main

import (
	"fmt"
	"math/rand"
	"time"

	. "main/linkedlist"
)

func main() {
	var ll LinkedList[int]

	seed := time.Now().UnixNano()
	random := rand.New(rand.NewSource(seed))

	// Insert Nodes
	const NODE_COUNT = 10
	const MAX = 1
	const MIN = 10
	for i := 0; i < NODE_COUNT; i++ {
		r := random.Intn(MIN-MAX+1) + MAX
		ll.Insert(&Node[int]{Data: r})
	}

	fmt.Println("Original\n--------")
	fmt.Println(ll.Details())
	fmt.Println(ll.String())

	ll1, ll2 := ll.SplitBy(ll.Mid())

	fmt.Println()
	fmt.Println("Split 1\n-------")
	fmt.Println(ll1.Details())
	fmt.Println(ll1.String())

	fmt.Println()
	fmt.Println("Split 2\n-------")
	fmt.Println(ll2.Details())
	fmt.Println(ll2.String())
}
