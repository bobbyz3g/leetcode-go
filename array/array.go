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

// ExistInRotatedSorted returns true if target is in nums,
// or false if it is not in nums.
//
// nums sorted in non-decreasing order (not necessarily with distinct values).
// and it is rotated at an unknown pivot index k (0 <= k < nums.length)
// such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1],
// nums[0], nums[1], ..., nums[k-1]] (0-indexed).
// For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5
// and become [4,5,6,6,7,0,1,2,4,4].
func ExistInRotatedSorted(nums []int, target int) bool {
	n := len(nums)
	f := func(i int) int {
		return nums[i] - target
	}
	cmp := func(i, j int) int {
		return nums[i] - nums[j]
	}
	index := SearchInRotatedSorted(n, f, cmp)
	return index >= 0
}

// SearchInRotatedSorted returns index if f(index) == 0.
// or -1 if it is not in array.
// f(index) return 0, if array[i] equal to target.
// cmp compares array[i] array[j], if array[i] equal array[j] returns 0,
// if array[i] less than array[j] return -1, else return 1.
func SearchInRotatedSorted(n int, f func(i int) int, cmp func(i, j int) int) int {
	if n == 0 {
		return -1
	}
	if n == 1 {
		if f(0) == 0 {
			return 0
		}
		return -1

	}

	l, r := 0, n-1
	for l <= r {
		mid := int((l + r) >> 1)
		if f(mid) == 0 {
			return mid
		}
		if cmp(l, mid) == 0 && cmp(mid, r) == 0 {
			l++
			r--
		} else if cmp(l, r) <= 0 {
			if f(l) <= 0 && f(mid) >= 0 {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if f(mid) < 0 && f(n-1) >= 0 {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}
