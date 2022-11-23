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
