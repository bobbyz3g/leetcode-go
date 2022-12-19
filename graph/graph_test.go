// Developed by Kaiser925 on 2022/12/19.
// Lasted modified 2022/12/19.
// Copyright (c) 2022.  All rights reserved
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graph

import "testing"

func TestValidPath(t *testing.T) {
	type args struct {
		n           int
		edges       [][]int
		source      int
		destination int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "exists",
			args: args{
				n: 3,
				edges: [][]int{
					{0, 1},
					{1, 2},
					{2, 0},
				},
				source:      0,
				destination: 2,
			},
			want: true,
		},
		{
			name: "not exists",
			args: args{
				n: 6,
				edges: [][]int{
					{0, 1},
					{0, 2},
					{3, 5},
					{5, 4},
					{4, 3},
				},
				source:      0,
				destination: 5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidPath(tt.args.n, tt.args.edges, tt.args.source, tt.args.destination); got != tt.want {
				t.Errorf("ValidPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
