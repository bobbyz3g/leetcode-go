// Developed by Kaiser925 on 2023/3/14.
// Lasted modified 2023/3/14.
// Copyright (c) 2023.  All rights reserved
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stack

import "testing"

func TestValidateStackSequences(t *testing.T) {
	type args struct {
		pushed []int
		popped []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				pushed: []int{1, 2, 3, 4, 5},
				popped: []int{4, 5, 3, 2, 1},
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				pushed: []int{1, 2, 3, 4, 5},
				popped: []int{4, 3, 5, 1, 2},
			},
			want: false,
		},
		{
			name: "case3",
			args: args{
				pushed: []int{1, 2, 3, 4, 5},
				popped: []int{1, 2, 3, 4, 5},
			},
			want: true,
		},
		{
			name: "case4",
			args: args{
				pushed: []int{2, 1, 0},
				popped: []int{1, 2, 0},
			},
			want: true,
		},
		{
			name: "case5",
			args: args{
				pushed: []int{1},
				popped: []int{1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateStackSequences(tt.args.pushed, tt.args.popped); got != tt.want {
				t.Errorf("ValidateStackSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
