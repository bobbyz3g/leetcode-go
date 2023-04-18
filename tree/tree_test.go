/*
 * Developed by Kaiser925 on 2022/11/22.
 * Lasted modified 2022/11/22.
 * Copyright (c) 2022.  All rights reserved
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1,null,2,3",
			args: args{
				&Node{
					Val: 1,
					Right: &Node{
						Val:  2,
						Left: &Node{Val: 3},
					},
				},
			},
			want: []int{1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InorderTraversal(tt.args.root); !assert.Equal(t, tt.want, got) {
				t.Errorf("InorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name string
		root *Node
		want int
	}{
		{
			name: "first",
			root: &Node{
				Left:  &Node{},
				Right: &Node{Left: &Node{}},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MaxDepth(tt.root), "MaxDepth(%v)", tt.root)
		})
	}
}

func TestFromSortedArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "1",
			args: args{
				nums: []int{1},
			},
			want: &Node{Val: 1},
		},
		{
			name: "[1,2]",
			args: args{
				nums: []int{1, 2},
			},
			want: &Node{Val: 1, Right: &Node{Val: 2}},
		},
		{
			name: "[-10,-3,0,5,9]",
			args: args{
				nums: []int{-10, -3, 0, 5, 9},
			},
			want: &Node{
				Left: &Node{
					Right: &Node{Val: -3},
					Val:   -10,
				},
				Val: 0,
				Right: &Node{
					Right: &Node{
						Val: 9,
					},
					Val: 5,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := FromSortedArray(tt.args.nums)
			equal := IsSameTree(bst, tt.want)
			assert.Equalf(t, true, equal, "FromSortedArray(%v)", tt.args.nums)
		})
	}
}

func TestIsSameTree(t *testing.T) {
	type args struct {
		p *Node
		q *Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "same",
			args: args{
				p: &Node{Val: 1},
				q: &Node{Val: 1},
			},
			want: true,
		},
		{
			name: "not same",
			args: args{
				p: &Node{Val: 1},
				q: &Node{Val: 1, Left: &Node{Val: 2}},
			},
			want: false,
		},
		{
			name: "both nil",
			args: args{
				p: nil,
				q: nil,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsSameTree(tt.args.p, tt.args.q), "IsSameTree(%v, %v)", tt.args.p, tt.args.q)
		})
	}
}

func TestIsBalanced(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil",
			args: args{
				root: nil,
			},
			want: true,
		},
		{
			name: "3,9,20,null,null,15,7",
			args: args{
				root: &Node{
					Right: &Node{
						Right: &Node{Val: 7},
						Val:   20,
						Left:  &Node{Val: 15},
					},
					Val:  3,
					Left: &Node{Val: 9},
				},
			},
			want: true,
		},
		{
			name: "1,2,2,3,3,null,null,4,4",
			args: args{
				root: &Node{
					Right: &Node{
						Val: 2,
					},
					Val: 1,
					Left: &Node{
						Right: &Node{
							Val: 3,
						},
						Val: 2,
						Left: &Node{
							Right: &Node{Val: 4},
							Val:   3,
							Left:  &Node{Val: 4},
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsBalanced(tt.args.root), "IsBalanced(%v)", tt.args.root)
		})
	}
}

func TestMinDepth(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[null]",
			args: args{
				root: nil,
			},
			want: 0,
		},
		{
			name: "[1,2,3,4]",
			args: args{
				root: &Node{
					Right: &Node{Val: 3, Left: &Node{Val: 4}},
					Val:   1,
					Left:  &Node{Val: 2},
				},
			},
			want: 2,
		},
		{
			name: "[2,null,3,null,4,null,5,null,6]",
			args: args{
				root: &Node{
					Val: 1,
					Left: &Node{
						Val: 2,
						Left: &Node{
							Val: 3,
							Left: &Node{
								Val:   4,
								Right: &Node{Val: 5},
							},
						},
					},
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MinDepth(tt.args.root), "MinDepth(%v)", tt.args.root)
		})
	}
}

func TestHasPathSum(t *testing.T) {
	type args struct {
		root      *Node
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil",
			args: args{
				root:      nil,
				targetSum: 0,
			},
			want: false,
		},
		{
			name: "[1,2,3]",
			args: args{
				root: &Node{
					Left:  &Node{Val: 2},
					Val:   1,
					Right: &Node{Val: 3},
				},
				targetSum: 5,
			},
			want: false,
		},
		{
			name: "[1,2]",
			args: args{
				root: &Node{
					Left: &Node{Val: 2},
					Val:  1,
				},
				targetSum: 1,
			},
			want: false,
		},
		{
			name: "[5,4,8,11,null,13,4,7,2,null,null,null,1]",
			args: args{
				root: &Node{
					Left: &Node{
						Left: &Node{
							Left:  &Node{Val: 7},
							Val:   11,
							Right: &Node{Val: 2},
						},
						Val: 4,
					},
					Val: 5,
					Right: &Node{
						Left: &Node{Val: 13},
						Val:  8,
						Right: &Node{
							Val:   4,
							Right: &Node{Val: 1},
						},
					},
				},
				targetSum: 22,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, HasPathSum(tt.args.root, tt.args.targetSum), "HasPathSum(%v, %v)", tt.args.root, tt.args.targetSum)
		})
	}
}

func TestPreorderTraversal(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1]",
			args: args{
				root: &Node{Val: 1},
			},
			want: []int{1},
		},
		{
			name: "[1,null,2,3]",
			args: args{
				root: &Node{
					Right: &Node{
						Val:  2,
						Left: &Node{Val: 3},
					},
					Val: 1,
				},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, PreorderTraversal(tt.args.root), "PreorderTraversal(%v)", tt.args.root)
		})
	}
}

func TestPostorderTraversal(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1]",
			args: args{
				root: &Node{Val: 1},
			},
			want: []int{1},
		},
		{
			name: "[1,null,2,3]",
			args: args{
				root: &Node{
					Right: &Node{
						Val:  2,
						Left: &Node{Val: 3},
					},
					Val: 1,
				},
			},
			want: []int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, PostorderTraversal(tt.args.root), "PostorderTraversal(%v)", tt.args.root)
		})
	}
}

func TestInvert(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "4,2,7,1,3,6,9",
			args: args{
				root: &Node{
					Right: &Node{
						Right: &Node{Val: 9},
						Val:   7,
						Left:  &Node{Val: 6},
					},
					Val: 4,
					Left: &Node{
						Right: &Node{Val: 3},
						Val:   2,
						Left:  &Node{Val: 1},
					},
				},
			},
			want: &Node{
				Left: &Node{
					Left:  &Node{Val: 9},
					Val:   7,
					Right: &Node{Val: 6},
				},
				Val: 4,
				Right: &Node{
					Left:  &Node{Val: 3},
					Val:   2,
					Right: &Node{Val: 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Invert(tt.args.root), "Invert(%v)", tt.args.root)
		})
	}
}

func TestMaxAncestorDiff(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				root: &Node{
					Right: &Node{
						Right: &Node{
							Val: 0,
							Left: &Node{
								Val: 3,
							},
						},
						Val: 2,
					},
					Val: 1,
				},
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				root: &Node{
					Right: &Node{
						Right: &Node{Val: 14,
							Left: &Node{Val: 13},
						},
						Val: 10},
					Val: 8,
					Left: &Node{
						Right: &Node{
							Right: &Node{Val: 7},
							Val:   6,
							Left:  &Node{Val: 4},
						},
						Val:  3,
						Left: &Node{Val: 1},
					},
				},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MaxAncestorDiff(tt.args.root), "MaxAncestorDiff(%v)", tt.args.root)
		})
	}
}
