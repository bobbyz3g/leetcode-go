package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func TestReverseList(t *testing.T) {
	test := NewList([]int{1, 2, 3, 4, 5})
	expected := NewList([]int{5, 4, 3, 2, 1})
	actual := reverseList(test)

	assert.Equal(t, true, expected.Equal(actual))
}
