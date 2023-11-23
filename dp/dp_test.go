package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{3, 7},
			want: 28,
		},
		{
			name: "case2",
			args: args{3, 2},
			want: 3,
		},
		{
			name: "case3",
			args: args{7, 3},
			want: 28,
		},
		{
			name: "case4",
			args: args{3, 3},
			want: 6,
		},
		{
			name: "case5",
			args: args{1, 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("UniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongestPalindrome(t *testing.T) {
	var tests5 = []struct {
		s   string
		out string
	}{
		{
			s:   "babad",
			out: "bab",
		},
		{
			s:   "cbbd",
			out: "bb",
		},
		{
			s:   "ac",
			out: "a",
		},
	}
	for _, v := range tests5 {
		assert.Equal(t, v.out, LongestPalindrome(v.s))
	}
}

func TestMinDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				word1: "horse",
				word2: "ros",
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				word1: "intention",
				word2: "execution",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MinDistance(tt.args.word1, tt.args.word2), "MinDistance(%v, %v)", tt.args.word1, tt.args.word2)
		})
	}
}

func TestNthUglyNumber(t *testing.T) {
	var tests264 = []struct {
		n   int
		num int
	}{
		{10, 12},
		{1, 1},
	}
	for _, v := range tests264 {
		assert.Equal(t, v.num, NthUglyNumber(v.n))
	}
}

func TestIntegerBreak(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				n: 2,
			},
			want: 1,
		},
		{
			name: "case2",
			args: args{
				n: 10,
			},
			want: 36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IntegerBreak(tt.args.n), "IntegerBreak(%v)", tt.args.n)
		})
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

func TestMaxSubarraySum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			want: 6,
		},
		{
			name: "case2",
			args: args{nums: []int{1}},
			want: 1,
		},
		{
			name: "case3",
			args: args{nums: []int{5, 4, -1, 7, 8}},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MaxSubarraySum(tt.args.nums), "MaxSubarraySum(%v)", tt.args.nums)
		})
	}
}

func TestMinPathSum(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{grid: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			}},
			want: 7,
		},
		{
			name: "case2",
			args: args{grid: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			}},
			want: 12,
		},
		{
			name: "case3",
			args: args{grid: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{1, 2, 1},
			}},
			want: 6,
		},
		{
			name: "case4",
			args: args{grid: [][]int{
				{1},
			}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MinPathSum(tt.args.grid), "MinPathSum(%v)", tt.args.grid)
		})
	}
}
