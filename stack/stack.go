// Developed by Kaiser925 on 2023/3/13.
// Lasted modified 2023/3/13.
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

import (
	"math"
)

type MinStack struct {
	arr      []int
	minStack []int
}

// NewMinStack returns new MinStack
func NewMinStack() *MinStack {
	return &MinStack{
		minStack: []int{math.MaxInt},
	}
}

// Push pushes new element to stack.
func (s *MinStack) Push(val int) {
	s.arr = append(s.arr, val)
	top := s.minStack[len(s.minStack)-1]
	s.minStack = append(s.minStack, min(top, val))
}

// Pop removes the element on the top of the stack.
func (s *MinStack) Pop() {
	if len(s.arr) == 0 {
		return
	}
	s.arr = s.arr[:len(s.arr)-1]
	s.minStack = s.minStack[:len(s.minStack)-1]
}

// Top returns top element in the stack.
func (s *MinStack) Top() int {
	if len(s.arr) == 0 {
		return 0
	}
	return s.arr[len(s.arr)-1]
}

// GetMin returns the minimum element in the stack.
func (s *MinStack) GetMin() int {
	return s.minStack[len(s.minStack)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ValidateStackSequences returns true if this could have been the result of
// a sequence of push and pop operations on an initially empty stack,
// or false otherwise. Elements are distinct in each sequence.
func ValidateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0, len(pushed))
	var j int
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	return len(stack) == 0
}

// DailyTemperatures returns an array answer such that
// answer[i] is the number of days you have to wait after
// the ith day to get a warmer temperature.
// If there is no future day for which this is possible,
// keep answer[i] == 0 instead.
func DailyTemperatures(temperatures []int) []int {
	// stack stores index of element inf temperatures
	stack := make([]int, 0, len(temperatures))
	res := make([]int, len(temperatures))
	for i, v := range temperatures {
		// stack is not empty, and top of stack is less than v.
		for n := len(stack); n != 0 && temperatures[stack[n-1]] < v; n = len(stack) {
			top := stack[n-1]
			stack = stack[:n-1]
			res[top] = i - top
		}
		stack = append(stack, i)
	}
	return res
}
