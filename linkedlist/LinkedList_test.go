package linkedlist

import "testing"

func TestLinkedList(t *testing.T) {
	t.Run("Check Size", func(t *testing.T) {
		cases := []struct {
			name     string
			input    *LinkedList[int]
			expected uint
		}{
			{
				name:     "Empty Node",
				input:    &LinkedList[int]{},
				expected: 0,
			},
			{
				name:     "Single Node",
				input:    (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}),
				expected: 1,
			},
			{
				name:     "Multiple Nodes",
				input:    (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}).Insert(&Node[int]{Data: 2}).Insert(&Node[int]{Data: 3}),
				expected: 3,
			},
			{
				name:     "Empty",
				input:    &LinkedList[int]{},
				expected: 0,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				actual := c.input.Size()
				if actual != c.expected {
					t.Errorf("Expected %v, got %v", c.expected, actual)
				}
			})
		}
	})

	t.Run("Check Equal", func(t *testing.T) {
		cases := []struct {
			name     string
			input1   *LinkedList[int]
			input2   *LinkedList[int]
			expected bool
		}{
			{
				name:     "Empty",
				input1:   &LinkedList[int]{},
				input2:   &LinkedList[int]{},
				expected: true,
			},
			{
				name:     "Single Node",
				input1:   (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}),
				input2:   (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}),
				expected: true,
			},
			{
				name:     "Multiple Nodes",
				input1:   (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}).Insert(&Node[int]{Data: 2}).Insert(&Node[int]{Data: 3}),
				input2:   (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}).Insert(&Node[int]{Data: 2}).Insert(&Node[int]{Data: 3}),
				expected: true,
			},
			{
				name:     "Different Size",
				input1:   (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}).Insert(&Node[int]{Data: 2}).Insert(&Node[int]{Data: 3}),
				input2:   (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}).Insert(&Node[int]{Data: 2}),
				expected: false,
			},
			{
				name:     "Different Data at Tail",
				input1:   (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}).Insert(&Node[int]{Data: 2}).Insert(&Node[int]{Data: 3}),
				input2:   (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}).Insert(&Node[int]{Data: 2}).Insert(&Node[int]{Data: 4}),
				expected: false,
			},
			{
				name:     "Different Data at Mid",
				input1:   (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}).Insert(&Node[int]{Data: 2}).Insert(&Node[int]{Data: 3}),
				input2:   (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}).Insert(&Node[int]{Data: 4}).Insert(&Node[int]{Data: 3}),
				expected: false,
			},
			{
				name:     "Different Data at Head",
				input1:   (&LinkedList[int]{}).Insert(&Node[int]{Data: 0}).Insert(&Node[int]{Data: 2}).Insert(&Node[int]{Data: 3}),
				input2:   (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}).Insert(&Node[int]{Data: 2}).Insert(&Node[int]{Data: 3}),
				expected: false,
			},
			{
				name:     "Different Data at other place",
				input1:   (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}).Insert(&Node[int]{Data: 6}).Insert(&Node[int]{Data: 2}).Insert(&Node[int]{Data: 3}),
				input2:   (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}).Insert(&Node[int]{Data: 5}).Insert(&Node[int]{Data: 2}).Insert(&Node[int]{Data: 3}),
				expected: false,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				actual := IsEqual(c.input1, c.input2)
				if actual != c.expected {
					t.Errorf("Expected %v, got %v", c.expected, actual)
				}
			})
		}
	})

	t.Run("Check Reverse In place", func(t *testing.T) {
		type Case struct {
			name     string
			actual   *LinkedList[int]
			expected *LinkedList[int]
		}

		cases := []Case{
			{
				name:     "Empty",
				actual:   &LinkedList[int]{},
				expected: &LinkedList[int]{},
			},
			{
				name:     "Single Node",
				actual:   (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}),
				expected: (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}),
			},
			{
				name:     "Odd Multiple Nodes",
				actual:   (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}).Insert(&Node[int]{Data: 2}).Insert(&Node[int]{Data: 3}),
				expected: (&LinkedList[int]{}).Insert(&Node[int]{Data: 3}).Insert(&Node[int]{Data: 2}).Insert(&Node[int]{Data: 1}),
			},
			{
				name:     "Even Multiple Nodes",
				actual:   (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}).Insert(&Node[int]{Data: 2}).Insert(&Node[int]{Data: 3}).Insert(&Node[int]{Data: 4}),
				expected: (&LinkedList[int]{}).Insert(&Node[int]{Data: 4}).Insert(&Node[int]{Data: 3}).Insert(&Node[int]{Data: 2}).Insert(&Node[int]{Data: 1}),
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				c.actual.Reverse()
				if !IsEqual(c.actual, c.expected) {
					t.Errorf("Expected %v, got %v", c.expected, c.actual)
				}
			})
		}
	})

	t.Run("Check Palindrome", func(t *testing.T) {
		cases := []struct {
			name     string
			input    *LinkedList[int]
			expected bool
		}{
			{
				name:     "Empty",
				input:    &LinkedList[int]{},
				expected: true,
			},
			{
				name:     "Single Node",
				input:    (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}),
				expected: true,
			},
			{
				name:     "Odd Multiple Nodes - Non-Palindrome",
				input:    (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}).Insert(&Node[int]{Data: 2}).Insert(&Node[int]{Data: 3}),
				expected: false,
			},
			{
				name:     "Even Multiple Nodes - Non-Palindrome",
				input:    (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}).Insert(&Node[int]{Data: 2}).Insert(&Node[int]{Data: 2}).Insert(&Node[int]{Data: 3}),
				expected: false,
			},
			{
				name:     "Odd Multiple Nodes - Palindrome",
				input:    (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}).Insert(&Node[int]{Data: 2}).Insert(&Node[int]{Data: 1}),
				expected: true,
			},
			{
				name:     "Even Multiple Nodes - Palindrome",
				input:    (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}).Insert(&Node[int]{Data: 2}).Insert(&Node[int]{Data: 2}).Insert(&Node[int]{Data: 1}),
				expected: true,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				actual := c.input.IsPalindrome()
				if actual != c.expected {
					t.Errorf("Expected %v, got %v", c.expected, actual)
				}
			})
		}
	})

	t.Run("Merge Lists", func(t *testing.T) {
		type Case struct {
			name     string
			input1   *LinkedList[int]
			input2   *LinkedList[int]
			expected uint
		}

		cases := []Case{
			{
				name:     "Single Node",
				input1:   (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}),
				input2:   (&LinkedList[int]{}).Insert(&Node[int]{Data: 2}),
				expected: 2,
			},
			{
				name:     "Multiple Nodes",
				input1:   (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}).Insert(&Node[int]{Data: 2}).Insert(&Node[int]{Data: 3}),
				input2:   (&LinkedList[int]{}).Insert(&Node[int]{Data: 4}).Insert(&Node[int]{Data: 5}).Insert(&Node[int]{Data: 6}),
				expected: 6,
			},
			{
				name:     "Different Size",
				input1:   (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}).Insert(&Node[int]{Data: 2}).Insert(&Node[int]{Data: 3}),
				input2:   (&LinkedList[int]{}).Insert(&Node[int]{Data: 4}).Insert(&Node[int]{Data: 5}),
				expected: 5,
			},
		}

		(func() {
			ll1 := (&LinkedList[int]{}).Insert(&Node[int]{Data: 4}).Insert(&Node[int]{Data: 5}).Insert(&Node[int]{Data: 6})
			llTmp := (&LinkedList[int]{}).Insert(&Node[int]{Data: 1}).Insert(&Node[int]{Data: 2}).Insert(&Node[int]{Data: 3})
			ll2, _ := llTmp.SplitBy(llTmp.Mid())

			cases = append(cases, Case{
				name:     "Splited",
				input1:   ll1,
				input2:   ll2,
				expected: 5,
			})
		})()

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				actual := c.input1.Merge(c.input2)
				if actual.Size() != c.expected {
					t.Errorf("Expected %v, got %v", c.expected, actual.Size())
				}
			})
		}
	})
}
