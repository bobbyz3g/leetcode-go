package algebra

import (
	"errors"
	"math"
	"sort"
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
		dp[i] = min(x2, x3, x5)
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
		result = max(result, right-left+1)
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

// MaximumUnits returns the maximum number of boxes that can be put on the truck.
// The numberOfBoxes[i] is the number of boxes of type i.
// The numberOfUnitsPerBox[i] is the number of units in each box of the type i
func MaximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	res := 0
	curSize := 0
outer:
	for i := 0; i < truckSize && i < len(boxTypes); i++ {
		for j := 0; j < boxTypes[i][0]; j++ {
			if curSize == truckSize {
				break outer
			}
			res += boxTypes[i][1]
			curSize++
		}
	}
	return res
}

// PlusOne increments the large integer by one and
// return the resulting array of digits.
func PlusOne(digits []int) []int {
	n := 1
	for i := len(digits) - 1; i >= 0; i-- {
		s := digits[i] + n
		digits[i] = s % 10
		n = s / 10
	}
	res := make([]int, len(digits)+1)
	for i := 0; i < len(digits); i++ {
		res[i+1] = digits[i]
	}
	res[0] = n

	if res[0] == 0 {
		return res[1:]
	}
	return res
}

// SqrtX returns  the square root of x rounded down to the nearest integer.
func SqrtX(x int) int {
	var xHalf = 0.5 * float64(x)
	var i = math.Float64bits(float64(x))
	i = 0x1FF7A3BEA91D9B1B + (i >> 1)
	var f = math.Float64frombits(i)
	f = f*0.5 + xHalf/f
	f = f*0.5 + xHalf/f
	f = f*0.5 + xHalf/f
	return int(f)
}

// ClimbStairs returns there are how many distinct ways
// can you climb to the top. Each time you can either
// climb 1 or 2 steps.
// f(x) = f(x-1) + f(x-2)
func ClimbStairs(x int) int {
	fx2, fx1, fx := 0, 0, 1
	for i := 1; i <= x; i++ {
		fx2 = fx1
		fx1 = fx
		fx = fx2 + fx1
	}
	return fx
}

// IsHappyNum returns true if a number n is happy.
// A happy number defined by follow process:
//
// Starting with any positive integer, replace the number by
// the sum of the squares of its digits.
// Repeat the process until the number equals 1 (where it will stay),
// or it loops endlessly in a cycle which does not include 1.
// Those numbers for which this process ends in 1 are happy.
func IsHappyNum(n int) bool {
	step := func(n int) int {
		sum := 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n = n / 10
		}
		return sum
	}
	slow, fast := n, step(n)
	for fast != 1 && slow != fast {
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1
}

// MaximumScore returns step tha make a, b c be zero.
// Every step can minus two of a, b, c.
func MaximumScore(a int, b int, c int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maxNum := max(a, max(b, c))
	s := a + b + c
	if s-maxNum < maxNum {
		return s - maxNum
	}
	return s / 2
}

// ReinitializePermutation returns the minimum non-zero number of
// operations you need to perform on perm to return the permutation to
// its initial value.
// Perm is a permutation perm of size n where perm[i] == i (0-indexed).
// In one operation, you will create a new array arr, and for each i:
// 1. If i % 2 == 0, then arr[i] = perm[i / 2]
// 2. If i % 2 == 1, then arr[i] = perm[n / 2 + (i - 1) / 2]
func ReinitializePermutation(n int) int {
	step := 0
	target := make([]int, n)
	for i := range target {
		target[i] = i
	}
	perm := append([]int(nil), target...)
	equal := func(a, b []int) bool {
		for i := range b {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	for {
		step++
		arr := make([]int, n)
		for i := range arr {
			if i%2 == 0 {
				arr[i] = perm[i/2]
			} else {
				arr[i] = perm[n/2+i/2]
			}
		}
		perm = arr
		if equal(target, perm) {
			break
		}
	}
	return step
}

func SmallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}

	cur, res := 0, 1

	for {
		cur = (10*cur + 1) % k
		if cur == 0 {
			return res
		}
		res++
	}
}

// ThreeSsumClosest returns the closest sum of
// three integers in nums to target
func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	best := math.MaxInt
	updateBest := func(cur int) {
		if abs(cur-target) < abs(best-target) {
			best = cur
		}
	}
	n := len(nums)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		p, q := i+1, n-1
		for p < q {
			s := nums[i] + nums[p] + nums[q]
			if s == target {
				return target
			}
			updateBest(s)
			if s > target {
				q1 := q - 1
				for p < q1 && nums[q1] == nums[q] {
					q1--
				}
				q = q1
			} else {
				p1 := p + 1
				for p1 < q && nums[p1] == nums[p] {
					p1++
				}
				p = p1
			}
		}
	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

// MyPow implements pow(x, n), which
// calculates x raised to the power n (i.e., xn).
func MyPow(x float64, n int) float64 {
	if n >= 0 {
		return pow(x, n)
	}
	return 1.0 / pow(x, -n)
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := pow(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}
