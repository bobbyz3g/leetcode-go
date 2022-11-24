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

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{
			input: []int{-4, -1, 0, 3, 10},
			want:  []int{0, 1, 9, 16, 100},
		},
		{
			input: []int{-7, -3, 2, 3, 11},
			want:  []int{4, 9, 9, 49, 121},
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, SortedSquares(tt.input))
	}
}

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{
			input: []int{0, 1, 0, 3, 12},
			want:  []int{1, 3, 12, 0, 0},
		},
		{
			input: []int{0},
			want:  []int{0},
		},
	}
	for _, tt := range tests {
		MoveZeroes(tt.input)
		assert.Equal(t, tt.want, tt.input)
	}
}

func TestTwoSumII(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{1, 2},
		},
		{
			nums:   []int{2, 3, 4},
			target: 6,
			want:   []int{1, 3},
		},
		{
			nums:   []int{-1, 0},
			target: -1,
			want:   []int{1, 2},
		},
	}
	for _, tt := range tests {
		assert.Equalf(t, tt.want, TwoSumII(tt.nums, tt.target), "test for %v %v", tt.nums, tt.target)
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		s    []byte
		want []byte
	}{
		{
			s:    []byte{'h', 'e', 'l', 'l', 'o'},
			want: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			s:    []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			want: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
		{
			s:    []byte{},
			want: []byte{},
		},
		{
			s:    []byte{'a'},
			want: []byte{'a'},
		},
	}

	for _, tt := range tests {
		ReverseString(tt.s)
		assert.Equal(t, tt.want, tt.s)
	}
}

func TestReverseWord(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			s:    "Let's take LeetCode contest",
			want: "s'teL ekat edoCteeL tsetnoc",
		},
		{
			s:    "God Ding",
			want: "doG gniD",
		},
		{
			s:    "God",
			want: "doG",
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, ReverseWord(tt.s))
	}
}

func TestFloodFill(t *testing.T) {
	tests := []struct {
		image    [][]int
		sr       int
		sc       int
		newColor int
		want     [][]int
	}{
		{
			image:    [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
			sr:       1,
			sc:       1,
			newColor: 2,
			want:     [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, FloodFill(tt.image, tt.sr, tt.sc, tt.newColor))
	}
}

func TestMaxAreaOfIsland(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{{0, 0, 0, 0, 0}},
			want: 0,
		},
		{
			grid: [][]int{
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, MaxAreaOfIsland(tt.grid))
	}
}

func TestMergeOrderArray(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1,2,3,0,0,0",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			want: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "1",
			args: args{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
			want: []int{1},
		},
		{
			name: "0, 1",
			args: args{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeOrderArray(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			assert.Equal(t, tt.args.nums1, tt.want)
		})
	}
}

func TestMaxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2,4,1",
			args: args{
				prices: []int{2, 4, 1},
			},
			want: 2,
		},
		{
			name: "7,6,4,3,1",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
		{
			name: "7,1,5,3,6,4",
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MaxProfit(tt.args.prices), "MaxProfit(%v)", tt.args.prices)
		})
	}
}
