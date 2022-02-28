package array

import (
	"sort"

	"github.com/Kaiser925/leetcode-go/algebra"
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
	dp := make([]int, l)
	dp[0] = 1
	result := 1

	for i := range nums {
		dp[i] = 1

		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = algebra.MaxInt(dp[i], dp[j]+1)
			}
		}

		result = algebra.MaxInt(result, dp[i])
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
		var t []int
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

// FindMaxLength returns the maximum length of
// a contiguous subarray with an equal number of 0 and 1.
func FindMaxLength(nums []int) (maxL int) {
	if len(nums) < 2 {
		return
	}
	mp := map[int]int{0: -1}
	cnt := 0
	for i, n := range nums {
		if n == 1 {
			cnt++
		} else {
			cnt--
		}
		if v, ok := mp[cnt]; ok {
			maxL = algebra.MaxInt(maxL, i-v)
		} else {
			mp[cnt] = i
		}
	}
	return
}

// BinarySearch searches the index of element that value is equal target in nums.
// If there is no element equal to target, returns -1.
// Note: array should be sorted in asc.
func BinarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1

	for left <= right {
		i := left + (right-left)/2
		if nums[i] == target {
			return i
		}

		if nums[i] < target {
			left = i + 1
		} else {
			right = i - 1
		}
	}
	return -1
}

// SortedSquares calculates the squares of slice element, and sort it.
func SortedSquares(nums []int) []int {
	n := len(nums)

	index := -1
	for i := 0; i < n && nums[i] < 0; i++ {
		index = i
	}
	res := make([]int, 0, n)

	neg, nonNeg := index, index+1
	for neg >= 0 || nonNeg < n {
		if neg < 0 {
			res = append(res, nums[nonNeg]*nums[nonNeg])
			nonNeg++
		} else if nonNeg == n {
			res = append(res, nums[neg]*nums[neg])
			neg--
		} else if nums[neg]*nums[neg] < nums[nonNeg]*nums[nonNeg] {
			res = append(res, nums[neg]*nums[neg])
			neg--
		} else {
			res = append(res, nums[nonNeg]*nums[nonNeg])
			nonNeg++
		}
	}

	return res
}

// MoveZeroes moves all 0's to the end of it while maintaining the
// relative order of the non-zero elements.
// link: https://leetcode-cn.com/problems/move-zeroes/
func MoveZeroes(nums []int) {
	n := len(nums)
	for i := range nums {
		if nums[i] == 0 {
			left := i

			for right := i + 1; right < n; right++ {
				if nums[right] != 0 {
					nums[left], nums[right] = nums[right], nums[left]
					break
				}
			}
		}
	}
}

// TwoSumII returns the indices of the two numbers, index1 and index2,
// added by one as an integer array [index1, index2] of length 2.
// link：https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted
func TwoSumII(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		s := numbers[left] + numbers[right]
		if s == target {
			return []int{left + 1, right + 1}
		} else if s < target {
			left++
		} else {
			right--
		}
	}
	return []int{-1, -1}
}

// ReverseString reverses a string.
// link: https://leetcode-cn.com/problems/reverse-string/
func ReverseString(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

// ReverseWord reverses every word in string.
// link: https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/
func ReverseWord(s string) string {
	b := make([]byte, 0, len(s))

	n := len(s)
	for i := 0; i < n; {
		l := i

		for i < n && s[i] != ' ' {
			i++
		}

		for r := i - 1; r >= l; r-- {
			b = append(b, s[r])
		}

		for i < n && s[i] == ' ' {
			i++
			b = append(b, ' ')
		}
	}

	return string(b)
}

var (
	dx = []int{1, 0, 0, -1}
	dy = []int{0, 1, -1, 0}
)

// FloodFill performs a flood fill on the image starting from the pixel image[sr][sc]
// link: https://leetcode-cn.com/problems/flood-fill/
func FloodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	curColor := image[sr][sc]

	if curColor != newColor {
		floodFillDfs(image, sr, sc, curColor, newColor)
	}

	return image
}

func floodFillDfs(image [][]int, x, y, color, newColor int) {
	if image[x][y] == color {
		image[x][y] = newColor
		for i := 0; i < 4; i++ {
			// get next position
			moveX, moveY := x+dx[i], y+dy[i]
			// if position is valid, continue make color.
			if moveX >= 0 && moveX < len(image) && moveY >= 0 && moveY < len(image[0]) {
				floodFillDfs(image, moveX, moveY, color, newColor)
			}
		}
	}
}
