package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{-1, head}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			tmp := cur.Next
			cur.Next = cur.Next.Next
			tmp.Next = nil
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

var tests83 = []struct {
	in  *ListNode
	out *ListNode
}{
	{
		NewList([]int{1, 2, 3, 3, 4, 4, 5}),
		NewList([]int{1, 2, 3, 4, 5}),
	},
	{
		NewList([]int{1, 1, 1, 2, 3}),
		NewList([]int{1, 2, 3}),
	},
	{
		NewList([]int{1, 1, 1}),
		NewList([]int{1}),
	},
	{
		NewList([]int{1, 2, 3}),
		NewList([]int{1, 2, 3}),
	},
}

func TestDeleteDuplicates(t *testing.T) {
	for _, v := range tests83 {
		res := deleteDuplicates(v.in)
		assert.True(t, res.Equal(v.out))
	}
}
