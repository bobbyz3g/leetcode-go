package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicates(t *testing.T) {
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
	for _, v := range tests83 {
		res := DeleteDuplicates(v.in)
		assert.True(t, res.Equal(v.out))
	}
}

func TestDeleteDuplicates2(t *testing.T) {
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

	for _, v := range tests82 {
		res := DeleteDuplicates2(v.in)
		assert.True(t, res.Equal(v.out))
	}
}

func TestReverseList(t *testing.T) {
	test := NewList([]int{1, 2, 3, 4, 5})
	expected := NewList([]int{5, 4, 3, 2, 1})
	actual := ReverseList(test)

	assert.Equal(t, true, expected.Equal(actual))
}

func TestReverseBetween(t *testing.T) {

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
	for _, v := range tests92 {
		assert.Equal(t, v.res, v.out.Equal(ReverseBetween(v.head, v.left, v.right)))
	}
}
