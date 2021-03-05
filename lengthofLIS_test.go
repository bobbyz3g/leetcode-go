package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func lengthOfLIS(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp := make([]int, l)
	dp[0] = 1
	result := 1

	for i := range nums {
		dp[i] = 1

		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		result = max(result, dp[i])
	}
	return result
}

var tests_300 = []struct {
	nums []int
	out  int
}{
	{
		nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
		out:  4,
	},
	{
		nums: []int{0, 1, 0, 3, 2, 3},
		out:  4,
	},
	{
		nums: []int{7, 7, 7, 7, 7, 7, 7},
		out:  1,
	},
	{
		nums: []int{},
		out:  0,
	},
}

func TestLengthOfLIS(t *testing.T) {
	for _, v := range tests_300 {
		assert.Equal(t, v.out, lengthOfLIS(v.nums))
	}
}
