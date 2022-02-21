package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayPairSum(t *testing.T) {
	var tests561 = []struct {
		input  []int
		output int
	}{
		{
			input:  []int{1, 4, 3, 2},
			output: 4,
		},
		{
			input:  []int{6, 2, 6, 5, 1, 2},
			output: 9,
		},
	}
	for _, v := range tests561 {
		assert.Equal(t, PairSum(v.input), v.output)
	}
}

func TestFindMaxLength(t *testing.T) {
	var tests525 = []struct {
		nums   []int
		output int
	}{
		{
			[]int{0, 1},
			2,
		},
		{
			[]int{0, 1, 0},
			2,
		},
		{
			[]int{0},
			0,
		},
	}

	for _, tt := range tests525 {
		assert.Equal(t, tt.output, FindMaxLength(tt.nums))
	}
}

func TestLengthOfLIS(t *testing.T) {
	var tests300 = []struct {
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
	for _, v := range tests300 {
		assert.Equal(t, v.out, LengthOfLIS(v.nums))
	}
}

func TestRemoveDuplicates2(t *testing.T) {

	var tests80 = []struct {
		in        []int
		out       int
		newNumber []int
	}{
		{
			[]int{1, 1, 1, 2, 2, 3},
			5,
			[]int{1, 1, 2, 2, 3},
		},
		{
			[]int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			7,
			[]int{0, 0, 1, 1, 2, 3, 3},
		},
		{
			[]int{1},
			1,
			[]int{1},
		},
		{
			[]int{},
			0,
			[]int{},
		},
		{
			[]int{1, 2, 3},
			3,
			[]int{1, 2, 3},
		},
	}
	for _, v := range tests80 {
		l := RemoveDuplicates2(v.in)
		assert.Equal(t, v.out, l)
		assert.Equal(t, v.newNumber, v.in[:l])
	}
}

func TestRemoveDuplicates1(t *testing.T) {

	var tests26 = []struct {
		in        []int
		out       int
		newNumber []int
	}{
		{
			[]int{1, 1, 2},
			2,
			[]int{1, 2},
		},
		{
			[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			5,
			[]int{0, 1, 2, 3, 4},
		},
	}

	for _, v := range tests26 {
		l := RemoveDuplicates1(v.in)
		assert.Equal(t, v.out, l)
		assert.Equal(t, v.newNumber, v.in[:l])
	}
}

func TestExistInRotatedSorted(t *testing.T) {
	var tests81 = []struct {
		nums   []int
		target int
		ok     bool
	}{
		{
			[]int{2, 5, 6, 0, 0, 1, 2},
			0,
			true,
		},
		{
			[]int{2, 5, 6, 0, 0, 1, 2},
			3,
			false,
		},
	}

	for _, v := range tests81 {
		assert.Equal(t, v.ok, ExistInRotatedSorted(v.nums, v.target))
	}
}

func TestSubsetsWithDup(t *testing.T) {
	var tests90 = []struct {
		nums []int
		res  [][]int
	}{
		{
			[]int{1, 2, 2},
			[][]int{
				nil,
				{1},
				{2},
				{1, 2},
				{2, 2},
				{1, 2, 2},
			},
		},
		{
			[]int{0},
			[][]int{
				nil,
				{0},
			},
		},
	}
	for _, v := range tests90 {
		assert.Equal(t, v.res, SubsetsWithDup(v.nums))
	}
}

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		nums  []int
		val   int
		index int
	}{
		{
			nums:  []int{-1, 0, 3, 5, 9, 12},
			val:   9,
			index: 4,
		},
		{
			nums:  []int{-1, 0, 3, 5, 9, 12},
			val:   2,
			index: -1,
		},
		{
			nums:  []int{},
			val:   2,
			index: -1,
		},
		{
			nums:  []int{5},
			val:   5,
			index: 0,
		},
		{
			nums:  []int{1},
			val:   2,
			index: -1,
		},
		{
			nums:  []int{1, 2},
			val:   2,
			index: 1,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.index, BinarySearch(tt.nums, tt.val))
	}
}
