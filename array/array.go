package array

import "sort"

// PairSum groups these integers into n pairs (a1, b1), (a2, b2), ..., (an, bn)
// such that the sum of min(ai, bi) for all i is maximized. Return the maximized sum.
// origin: leetcode 561
func PairSum(nums []int) int {
	sort.Ints(nums)
	result := 0
	for i := 0; i < len(nums); i += 2 {
		result += nums[i]
	}
	return result
}

// LengthOfLIS returns the length of the longest strictly increasing subsequence.
func LengthOfLIS(nums []int) int {
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
