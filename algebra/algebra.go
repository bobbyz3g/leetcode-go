package algebra

import (
	"errors"
	"strings"
)

// ErrOutOfBound is a error that input number is not less than 4000 and greater than 0
var ErrOutOfBound = errors.New("n must be less than 4000 and greater than 0")

var (
	factors   = []int{2, 3, 5}
	thousands = []string{"", "M", "MM", "MMM"}
	hundreds  = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens      = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones      = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
)

// IntPair represents a pair of two int nums.
type IntPair struct{ Left, Right int }

// IsUgly returns true if n is a ugly number.
// Ugly number is a positive number whose prime factors only include 2, 3, and/or 5
func IsUgly(n int) bool {
	if n <= 0 {
		return false
	}

	for _, v := range factors {
		for n%v == 0 {
			n /= v
		}
	}

	return n == 1
}

// NthUglyNumber returns the nth ugly number.
// Ugly number is a positive number whose prime factors only include 2, 3, and/or 5
func NthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1

	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = MinInt(MinInt(x2, x3), x5)
		if dp[i] == x2 {
			p2++
		}
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
	}
	return dp[n]
}

// MinInt returns the minimum number.
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxInt returns the maximum number.
func MaxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// GetRow returns the nth (0-indexed) row of the Pascal's triangle.
func GetRow(n int) []int {
	l := n + 1
	result := make([]int, l)
	result[0] = 1

	for i := 0; i < n; i++ {
		pre := result[0]
		for j := 1; j < l; j++ {
			cur := result[j]
			result[j] = cur + pre
			pre = cur
		}
	}
	return result
}

// FindDisappearedNumbers returns an array of all the integers
// in the range [1, n] that do not appear in nums.
func FindDisappearedNumbers(nums []int) []int {
	result := make([]int, 0, 1)
	l := len(nums)

	// local array hash.
	for _, v := range nums {
		index := (v - 1) % l
		nums[index] = nums[index] + l
	}

	for i, v := range nums {
		if v <= l {
			result = append(result, i+1)
		}
	}

	return result
}

// LongestOnes return the maximum number of consecutive 1's
// in the array if you can flip at most k 0's
//
// origin：https://leetcode-cn.com/problems/max-consecutive-ones-iii
func LongestOnes(A []int, K int) int {
	var result, left, leftZeroNum, rightZeroNum int
	for right, v := range A {
		rightZeroNum = 1 - v + rightZeroNum
		for leftZeroNum < rightZeroNum-K {
			leftZeroNum = 1 - A[left] + leftZeroNum
			left++
		}
		result = MaxInt(result, right-left+1)
	}
	return result
}

// IsValidSerialization returns true if it is a correct preorder traversal serialization of a binary tree.
func IsValidSerialization(preorder string) bool {
	diff := 1
	for _, c := range strings.Split(preorder, ",") {
		diff--
		if diff < 0 {
			return false
		}
		if c != "#" {
			diff += 2
		}
	}
	return diff == 0
}

// SearchInsert return the index if the target is found.
// If not, return the index where it would be if it were inserted in order
func SearchInsert(nums []int, target int) int {
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

// IToR converts unsigned integer to a roman numeral.
// 1 <= num <= 3999
func IToR(n int) (string, error) {
	if n < 1 || n > 3000 {
		return "", ErrOutOfBound
	}
	return thousands[n/1000] + hundreds[n%1000/100] + tens[n%100/10] + ones[n%10], nil
}
