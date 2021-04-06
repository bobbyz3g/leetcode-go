package array

import (
	"sort"
)

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

// SubsetsWithDup returns all possible subsets (the power set).
// The result is not contain duplicate subsets.
func SubsetsWithDup(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	n := len(nums)
outer:
	for mask := 0; mask < 1<<n; mask++ {
		t := []int{}
		for i, v := range nums {
			if mask>>i&1 > 0 {
				if i > 0 && mask>>(i-1)&1 == 0 && v == nums[i-1] {
					continue outer
				}
				t = append(t, v)
			}
		}
		res = append(res, append([]int(nil), t...))
	}
	return res
}

// RemoveDuplicates1 removes the duplicates in-place,
// such that each element appears only once and returns the new length.
// Given nums is sorted.
func RemoveDuplicates1(nums []int) int {
	return RemoveDuplicates(nums, 1)
}

// RemoveDuplicates2 removes the duplicates in-place such that duplicates appeared at most twice
// and returns the new length slice.
// Given nums is sorted.
func RemoveDuplicates2(nums []int) int {
	return RemoveDuplicates(nums, 2)
}

// RemoveDuplicates removes the duplicates in-place,
// such that duplicates appeared at most appearNum and returns the new length.
// Given nums is sorted.
func RemoveDuplicates(nums []int, appearNum int) int {
	if len(nums) < appearNum || appearNum < 1 {
		return len(nums)
	}
	slow := appearNum - 1
	prev := appearNum - 1

	for fast := appearNum; fast < len(nums); fast++ {
		if nums[slow-prev] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}
