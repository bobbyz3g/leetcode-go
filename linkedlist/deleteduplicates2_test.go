package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{-1, head}

	cur := dummy
	next := cur.Next

	for next != nil && next.Next != nil {
		if next.Val == next.Next.Val {
			val := next.Val
			for next != nil && next.Val == val {
				next = next.Next
				cur.Next = next
			}
		} else {
			cur = next
			next = cur.Next
		}
	}

	return dummy.Next
}

var tests82 = []struct {
	in  *ListNode
	out *ListNode
}{
	{
		NewList([]int{1, 2, 3, 3, 4, 4, 5}),
		NewList([]int{1, 2, 5}),
	},
	{
		NewList([]int{1, 1, 1, 2, 3}),
		NewList([]int{2, 3}),
	},
	{
		NewList([]int{1, 1, 1}),
		NewList([]int{}),
	},
	{
		NewList([]int{1, 2, 3}),
		NewList([]int{1, 2, 3}),
	},
}

func TestDeleteDuplicates2(t *testing.T) {
	for _, v := range tests82 {
		res := deleteDuplicates2(v.in)
		assert.True(t, res.Equal(v.out))
	}
}
