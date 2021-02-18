package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func searchInsert(nums []int, target int) int {
	l := len(nums)
	result := l
	left, right := 0, l-1
	for left <= right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return result
}

var tests_995 = []struct {
	nums   []int
	target int
	output int
}{
	{
		nums:   []int{1, 3, 5, 6},
		target: 5,
		output: 2,
	},
	{
		nums:   []int{1, 3, 5, 6},
		target: 2,
		output: 1,
	},
	{
		nums:   []int{1, 3, 5, 6},
		target: 7,
		output: 4,
	},
	{
		nums:   []int{1, 3, 5, 6},
		target: 0,
		output: 0,
	},
}

func TestMinKBitFlips(t *testing.T) {
	for _, v := range tests_995 {
		assert.Equal(t, searchInsert(v.nums, v.target), v.output)
	}
}
