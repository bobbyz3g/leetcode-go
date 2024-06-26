package array

import (
	"math"
	"sort"
	"strings"
	"sync"
	"time"
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
func RemoveDuplicates1[T comparable](nums []T) int {
	return RemoveDuplicates(nums, 1)
}

// RemoveDuplicates2 removes the duplicates in-place such that duplicates appeared at most twice
// and returns the new length slice.
// Given nums is sorted.
func RemoveDuplicates2[T comparable](nums []T) int {
	return RemoveDuplicates(nums, 2)
}

// RemoveDuplicates removes the duplicates in-place,
// such that duplicates appeared at most appearNum and returns the new length.
// Given nums is sorted.
func RemoveDuplicates[T comparable](nums []T, appearNum int) int {
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
			maxL = max(maxL, i-v)
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
	match := func(a int) bool {
		return a == 0
	}
	MoveToEnd(nums, match)
}

// MoveToEnd moves match element to end.
func MoveToEnd[T any](arr []T, match func(T) bool) {
	n := len(arr)
	for i := range arr {
		if match(arr[i]) {
			left := i
			for right := i + 1; right < n; right++ {
				if !match(arr[right]) {
					arr[left], arr[right] = arr[right], arr[left]
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

// MaxAreaOfIsland returns the maximum area of an island in grid.
// If there is no island, return 0.
// An island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.)
//
// link：https://leetcode-cn.com/problems/max-area-of-island
func MaxAreaOfIsland(grid [][]int) int {
	res := 0
	for i, s := range grid {
		for j := range s {
			res = max(maxAreaOfIslandDfs(grid, i, j), res)
		}
	}
	return res
}

func maxAreaOfIslandDfs(grid [][]int, i, j int) int {
	if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) || grid[i][j] != 1 {
		return 0
	}
	grid[i][j] = 0
	res := 1
	for _, pos := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		nextI, nextJ := i+pos[0], j+pos[1]
		res += maxAreaOfIslandDfs(grid, nextI, nextJ)
	}
	return res
}

// MergeOrderArray merges nums1 and nums2 into a single array
// sorted in non-decreasing order.
// The final sorted array should not be returned by the function,
// but instead be stored inside the array nums1.
// To accommodate this, nums1 has a length of m + n,
// where the first m elements denote the elements that should be merged,
// and the last n elements are set to 0 and should be ignored.
// nums2 has a length of n.
func MergeOrderArray(nums1 []int, m int, nums2 []int, n int) {
	a, b := m-1, n-1
	for i := m + n - 1; b >= 0; i-- {
		if a >= 0 && nums1[a] > nums2[b] {
			nums1[i] = nums1[a]
			a--
		} else {
			nums1[i] = nums2[b]
			b--
		}
	}
}

// MaxProfit returns the maximum profit you can achieve from this transaction.
// If you cannot achieve any profit, return 0.
func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	maxProfit := 0
	minVal := math.MaxInt
	for _, v := range prices {
		if v < minVal {
			minVal = v
		} else if profit := v - minVal; profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}

// MajorityElement returns the majority element which appears more than ⌊n / 2⌋ times.
func MajorityElement[T any](s []T, eq func(a, b T) bool) T {
	winner := s[0]
	count := 0
	for i := 0; i < len(s); i++ {
		if eq(winner, s[i]) {
			count++
		} else if count == 0 {
			winner = s[i]
			count++
		} else {
			count--
		}
	}
	return winner
}

// PascalTriangle returns the first numRows of Pascal's triangle.
//
// In Pascal's triangle, each number is the sum
// of the two numbers directly above it
func PascalTriangle(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	rows := make([][]int, numRows)
	rows[0] = []int{1}

	for i := 1; i < numRows; i++ {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1
		for j := 1; j < i; j++ {
			row[j] = rows[i-1][j] + rows[i-1][j-1]
		}
		rows[i] = row
	}
	return rows
}

// ContainsDuplicate reports whether any value appears at
// least twice in the array
func ContainsDuplicate(nums []int) bool {
	m := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		_, ok := m[v]
		if ok {
			return true
		}
		m[v] = struct{}{}
	}
	return false
}

// ContainsNearbyDuplicate returns true if there are two distinct indices i
// and j in the array such that nums[i] == nums[j] and abs(i - j) <= k
func ContainsNearbyDuplicate(nums []int, k int) bool {
	indexes := make(map[int]int, len(nums))
	for i, v := range nums {
		prev, ok := indexes[v]
		if !ok {
			indexes[v] = i
			continue
		}
		if i-prev <= k {
			return true
		}
		indexes[v] = i
	}
	return false
}

// MinimumSize returns the minimum possible penalty after
// performing the operations. Penalty is the maximum number of
// balls in a bag.
func MinimumSize(nums []int, maxOperations int) int {
	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}

	return sort.Search(maxVal, func(i int) bool {
		if i == 0 {
			return false
		}
		ops := 0
		for _, x := range nums {
			ops += (x - 1) / i
		}
		return ops <= maxOperations
	})
}

const hourMinute = "15:04"

func AlertNames(keyName []string, keyTime []string) []string {
	timeMap := make(map[string][]time.Time)
	for i, name := range keyName {
		t, _ := time.Parse(hourMinute, keyTime[i])
		timeMap[name] = append(timeMap[name], t)
	}

	res := make([]string, 0, len(timeMap))
	for name, times := range timeMap {
		sort.Slice(times, func(i, j int) bool {
			return times[i].Before(times[j])
		})

		for i, t := range times[2:] {
			if t.Sub(times[i]).Hours() <= 1 {
				res = append(res, name)
				break
			}
		}
	}
	sort.Strings(res)
	return res
}

// RemoveSubfolders returns the folders after removing all sub-folders in those folders.
func RemoveSubfolders(folder []string) []string {
	sort.Strings(folder)
	res := make([]string, 0, len(folder))
	res = append(res, folder[0])
	for _, f := range folder[1:] {
		last := res[len(res)-1]
		if !strings.HasPrefix(f, last) || f[len(last)] != '/' {
			res = append(res, f)
		}
	}
	return res
}

// FourSum returns an array of all the unique quadruplets
// [nums[a], nums[b], nums[c], nums[d]] such that
// nums[a] + nums[b] + nums[c] + nums[d] == target, and
// a, b, c and d are distinct.
func FourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := make([][]int, 0, 5)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return res
}

func LongestWPI(hours []int) int {
	s := 0
	pos := make(map[int]int)
	res := 0
	for i, v := range hours {
		if v > 8 {
			s++
		} else {
			s--
		}

		if s > 0 {
			res = i + 1
		} else if j, ok := pos[s-1]; ok {
			res = max(res, i-j)
		}
		if _, ok := pos[s]; !ok {
			pos[s] = i
		}
	}
	return res
}
func NumberOfPairs(nums []int) []int {
	res := []int{0, len(nums)}
	m := make(map[int]struct{})
	for _, v := range nums {
		_, ok := m[v]
		if ok {
			delete(m, v)
			res[1] -= 2
			res[0] += 1
		} else {
			m[v] = struct{}{}
		}
	}
	return res
}

func GenerateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	m := make(map[string]struct{})
	for _, v := range GenerateParenthesis(n - 1) {
		for i := 0; i < 2*(n-1); i++ {
			k := v[0:i] + "()" + v[i:]
			m[k] = struct{}{}
		}
	}
	res := make([]string, 0, len(m))
	for k := range m {
		res = append(res, k)
	}
	sort.Strings(res)
	return res
}

// Permute returns all the possible permutations of nums.
func Permute(nums []int) [][]int {
	n := len(nums)
	visited := make([]bool, n)
	res := make([][]int, 0, n)
	permute(nums, visited, make([]int, n), 0, &res)
	return res
}

func permute(nums []int, visited []bool, walker []int, n int, res *[][]int) {
	if n == len(nums) {
		tmp := make([]int, n)
		copy(tmp, walker)
		*res = append(*res, tmp)
		return
	}
	for i, v := range nums {
		if visited[i] {
			continue
		}
		visited[i] = true
		walker[n] = v
		n++
		permute(nums, visited, walker, n, res)
		n--
		visited[i] = false
	}
}

// SearchRange returns starting and ending position of given target.
// Array was sorted in non-decreasing order.
func SearchRange(nums []int, target int) []int {
	left := sort.SearchInts(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	right := sort.SearchInts(nums, target+1)
	return []int{left, right - 1}
}

func MovesToMakeZigzag(nums []int) int {
	s := [2]int{}
	for i, v := range nums {
		left, right := math.MaxInt, math.MaxInt
		if i > 0 {
			left = nums[i-1]
		}
		if i < len(nums)-1 {
			right = nums[i+1]
		}
		s[i%2] += max(v-min(left, right)+1, 0)

	}
	return min(s[0], s[1])
}

func MergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	valWeights := make(map[int]int, len(items1)+len(items2))
	for _, v := range items1 {
		valWeights[v[0]] += v[1]
	}

	for _, v := range items2 {
		valWeights[v[0]] += v[1]
	}

	res := make([][]int, 0, len(valWeights))
	for k, v := range valWeights {
		res = append(res, []int{k, v})
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})
	return res
}

func MostFrequentEven(nums []int) int {
	count := make(map[int]int)

	for _, v := range nums {
		if v%2 == 1 {
			continue
		}
		count[v] += 1
	}
	if len(count) == 0 {
		return -1
	}

	res, ct := -1, 0
	for k, v := range count {
		if res == -1 || ct < v || ct == v && k < res {
			res = k
			ct = v
		}
	}
	return res
}

// MergeOverlapping merges all overlapping intervals,
// and returns an array of the non-overlapping intervals that
// cover all the intervals in the input.
func MergeOverlapping(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0, len(intervals))
	s, e := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > e {
			res = append(res, []int{s, e})
			s, e = intervals[i][0], intervals[i][1]
			continue
		}
		e = max(e, intervals[i][1])
	}
	res = append(res, []int{s, e})
	return res
}

func EliminateMaximum(dist []int, speed []int) int {
	n := len(dist)
	arrive := make([]int, n)
	for i := range dist {
		arrive[i] = (dist[i]-1)/speed[i] + 1
	}
	sort.Ints(arrive)
	for i := 0; i < n; i++ {
		if arrive[i] <= i {
			return i
		}
	}
	return n
}

func NextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	i, j := n-2, n-1
	for ; i >= 0 && nums[i] >= nums[j]; i, j = i-1, j-1 {
	}

	k := n - 1
	if i >= 0 {
		for nums[i] >= nums[k] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}

	for i, j := j, n-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func RepairCars(ranks []int, cars int) int64 {
	minR := ranks[0]
	for _, v := range ranks {
		if v < minR {
			minR = v
		}
	}

	return int64(sort.Search(minR*cars*cars, func(i int) bool {
		s := 0
		for _, v := range ranks {
			s += int(math.Sqrt(float64(i / v)))
		}
		return s >= cars
	}))
}

// InsertAndMergeOverlapping inserts newInterval into intervals,
// than merge overlapping interval.
func InsertAndMergeOverlapping(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	n := len(intervals)
	i := 0
	for i < n && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}
	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	res = append(res, newInterval)
	for i < n {
		res = append(res, intervals[i])
		i++
	}
	return res
}

func Jump(nums []int) int {
	end := 0
	next := 0
	steps := 0
	for i := 0; i < len(nums)-1; i++ {
		next = max(next, i+nums[i])
		if i == end {
			end = next
			steps++
		}
	}
	return steps
}

func CanJump(nums []int) bool {
	j := 0
	for i := 0; i < len(nums); i++ {
		if i > j {
			return false
		}
		j = max(j, nums[i]+i)
	}
	return true
}

func Intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0, len(nums1))
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		sort.Ints(nums1)
	}()
	go func() {
		defer wg.Done()
		sort.Ints(nums2)
	}()
	wg.Wait()

	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		a, b := nums1[i], nums2[j]
		if a == b {
			if len(res) == 0 || a > res[len(res)-1] {
				res = append(res, a)
			}
			i++
			j++
		} else if a < b {
			i++
		} else {
			j++
		}
	}

	return res
}

// SearchRotated returns index of target. If
// target is not in array, it will return -1.
// Nums is possibly rotated at an unknown pivot
// index k (1 <= k < nums.length) such that
// the resulting array is [nums[k], nums[k+1], ...,
// nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
func SearchRotated(nums []int, target int) int {
	n := len(nums)
	switch n {
	case 0:
		return -1
	case 1:
		if nums[0] == target {
			return 0
		}
		return -1
	}

	l, r := 0, n-1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return mid
		}
		// left is ordered
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
			continue
		}

		// right is ordered
		if nums[mid] < target && target <= nums[n-1] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func SortColors(nums []int) {
	rIdx, wIdx := 0, 0
	for i, c := range nums {
		if c == 0 {
			nums[i], nums[rIdx] = nums[rIdx], nums[i]
			if rIdx < wIdx {
				nums[i], nums[wIdx] = nums[wIdx], nums[i]
			}
			rIdx++
			wIdx++
		} else if c == 1 {
			nums[i], nums[wIdx] = nums[wIdx], nums[i]
			wIdx++
		}
	}
}

// MaxProfit2 calculates the maximum profit that can be achieved
// by buying and selling stocks from the given prices array.
// It uses dynamic programming to compute the value.
// The function returns the maximum profit.
// origin: leetcode 122
func MaxProfit2(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	dp := make([][2]int, n)
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func AddMinimum(word string) int {
	dp := make([]int, len(word)+1)
	for i := 1; i <= len(word); i++ {
		dp[i] = dp[i-1] + 2
		if i > 1 && word[i-1] > word[i-2] {
			dp[i] = dp[i-1] - 1
		}
	}
	return dp[len(word)]
}

// Subset generates all possible subsets of a given slice.
// It returns a slice of slices, where each element is a subset of the input slice.
// The order of the elements in the resulting subsets is not guaranteed.
// The function uses the binary counting method to generate the subsets.
// It iterates through all the possible binary numbers of length n,
// where n is the length of the input slice.
// For each binary number, the function checks which bits are set to 1,
// and includes the corresponding elements from the input slice in the subset.
// The function then appends the generated subset to the list of subsets.
// Finally, it returns the list of subsets.
//
// For example, given the input slice [1, 2, 3],
// the function will generate the following subsets:
// [], [1], [2], [1, 2], [3], [1, 3], [2, 3], [1, 2, 3]
//
// The function has a type parameter T which represents the type of elements in the input slice.
// The function signature is `func Subset[T any](a []T) [][]T`.
//
// Usage example:
// input := []int{1, 2, 3}
// subsets := Subset(input)
// // subsets = [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}
func Subset[T any](a []T) [][]T {
	count := 1 << len(a)
	subset := make([][]T, 0, len(a))
	for i := 0; i < count; i++ {
		sub := make([]T, 0)
		for j := 0; j < len(a); j++ {
			if (i>>j)&1 == 1 {
				sub = append(sub, a[j])
			}
		}
		subset = append(subset, sub)
	}
	return subset
}

// MaximumNumberOfStringPairs counts the number of pairs of strings in the given slice
// where swapping the first and second characters of a pair creates a duplicate pair.
// It returns the count of such pairs.
func MaximumNumberOfStringPairs(words []string) int {
	var count int
	exists := [26][26]bool{}
	for _, s := range words {
		x, y := s[0]-'a', s[1]-'a'
		if exists[y][x] {
			count++
		} else {
			exists[x][y] = true
		}
	}
	return count
}

// MinimumRemoval calculates the minimum number of removals required to obtain an array
// where each element is its index multiplied by the total sum of the array.
// It takes in an integer array `beans` and returns the minimum number of removals as an int64.
// The function first sorts the array in ascending order using the `sort.Ints` function.
// It calculates the total sum of all elements in the array using a loop and stores it in `total`.
// Then, it iterates over each element in the sorted array and calculates the number of removals
// required by subtracting the current element multiplied by the remaining number of elements from the total sum.
// The function keeps track of the minimum number of removals in the `removal` variable.
// Finally, it returns the value of `removal` as the result.
//
// origin: leetcode problem 2171
func MinimumRemoval(beans []int) int64 {
	sort.Ints(beans)
	var total int64
	for _, v := range beans {
		total += int64(v)
	}
	removal := total
	n := len(beans)
	for i, bean := range beans {
		removal = min(removal, total-int64(bean)*int64(n-i))
	}
	return removal
}

func GroupAnagrams(strs []string) [][]string {
	key := func(s string) string {
		b := []byte(s)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		return string(b)
	}

	m := make(map[string][]string)

	for _, s := range strs {
		k := key(s)
		m[k] = append(m[k], s)
	}

	res := make([][]string, 0, len(m))

	for _, s := range m {
		res = append(res, s)
	}

	return res
}
