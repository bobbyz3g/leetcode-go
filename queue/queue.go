// Developed by Kaiser925 on 2023/3/20.
// Lasted modified 2023/3/20.
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

package queue

import (
	"container/heap"
	"sort"
)

type PriorityQueue struct {
	a []int
	sort.IntSlice
}

func NewPriorityQueue(elems []int, size int) *PriorityQueue {
	q := &PriorityQueue{
		a:        elems,
		IntSlice: make([]int, 0, size),
	}
	for i := 0; i < size; i++ {
		q.IntSlice = append(q.IntSlice, i)
	}
	return q
}

func (p *PriorityQueue) Less(i, j int) bool {
	return p.a[p.IntSlice[i]] > p.a[p.IntSlice[j]]
}

func (p *PriorityQueue) Push(v interface{}) {
	p.IntSlice = append(p.IntSlice, v.(int))
}

func (p *PriorityQueue) Pop() interface{} {
	v := p.IntSlice[len(p.IntSlice)-1]
	p.IntSlice = p.IntSlice[:len(p.IntSlice)-1]
	return v
}

// MaxSlidingWindow return max value in sliding window which size is k.
// Each time the sliding window moves right by one position
func MaxSlidingWindow(nums []int, k int) []int {
	q := NewPriorityQueue(nums, k)
	heap.Init(q)

	n := len(nums)
	res := make([]int, 0, n-k+1)
	res = append(res, nums[q.IntSlice[0]])

	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		res = append(res, nums[q.IntSlice[0]])
	}

	return res
}
