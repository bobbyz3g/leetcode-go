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

func TestHasCycle(t *testing.T) {
	type args struct {
		nodes []*ListNode
		pos   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				nodes: []*ListNode{{Val: 1}},
				pos:   -1,
			},
			want: false,
		},
		{
			name: "3,2,0,-4",
			args: args{
				nodes: []*ListNode{{Val: 3}, {Val: 2}, {Val: 0}, {Val: -4}},
				pos:   1,
			},
			want: true,
		},
		{
			name: "1,2",
			args: args{
				nodes: []*ListNode{{Val: 1}, {Val: 2}},
				pos:   0,
			},
			want: true,
		}, {
			name: "1,2",
			args: args{
				nodes: []*ListNode{{Val: 1}, {Val: 2}},
				pos:   -1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := tt.args.nodes[0]
			n := len(tt.args.nodes)
			node := head
			if tt.args.pos >= 0 {
				tt.args.nodes[n-1].Next = tt.args.nodes[tt.args.pos]
			}
			for i := 1; i < n; i++ {
				node.Next = tt.args.nodes[i]
				node = node.Next
			}
			assert.Equalf(t, tt.want, HasCycle(head), "")
		})
	}
}

func TestRemoveElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "empty",
			args: args{
				head: nil,
				val:  0,
			},
			want: nil,
		},
		{
			name: "remove all",
			args: args{
				head: NewList([]int{7, 7, 7, 7}),
				val:  7,
			},
			want: nil,
		},
		{
			name: "remove last",
			args: args{
				head: NewList([]int{7, 7, 7, 8}),
				val:  8,
			},
			want: NewList([]int{7, 7, 7}),
		},
		{
			name: "remove first",
			args: args{
				head: NewList([]int{8, 7, 7, 7}),
				val:  8,
			},
			want: NewList([]int{7, 7, 7}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, RemoveElements(tt.args.head, tt.args.val), "RemoveElements(%v, %v)", tt.args.head, tt.args.val)
		})
	}
}

func TestMergeInBetween(t *testing.T) {
	type args struct {
		list1 *ListNode
		a     int
		b     int
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case 1",
			args: args{
				list1: NewList([]int{0, 1, 2, 3, 4, 5}),
				a:     3,
				b:     4,
				list2: NewList([]int{1000000, 1000001, 1000002}),
			},
			want: NewList([]int{0, 1, 2, 1000000, 1000001, 1000002, 5}),
		},
		{
			name: "case 2",
			args: args{
				list1: NewList([]int{0, 1, 2, 3, 4, 5, 6}),
				a:     2,
				b:     5,
				list2: NewList([]int{1000000, 1000001, 1000002, 1000003, 1000004}),
			},
			want: NewList([]int{0, 1, 1000000, 1000001, 1000002, 1000003, 1000004, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MergeInBetween(tt.args.list1, tt.args.a, tt.args.b, tt.args.list2), "MergeInBetween(%v, %v, %v, %v)", tt.args.list1, tt.args.a, tt.args.b, tt.args.list2)
		})
	}
}
