package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummy.Next
}

var tests92 = []struct {
	head  *ListNode
	left  int
	right int
	out   *ListNode
	res   bool
}{
	{
		NewList([]int{1, 2, 3, 4, 5}),
		2,
		4,
		NewList([]int{1, 4, 3, 2, 5}),
		true,
	},
	{
		NewList([]int{3, 5}),
		1,
		2,
		NewList([]int{5, 3}),
		true,
	},
	{
		NewList([]int{5}),
		1,
		1,
		NewList([]int{5}),
		true,
	},
}

func TestReverseBetween(t *testing.T) {
	for _, v := range tests92 {
		assert.Equal(t, v.res, v.out.Equal(reverseBetween(v.head, v.left, v.right)))
	}
}
