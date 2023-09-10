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

func TestMajorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "3,2,3",
			args: args{
				nums: []int{3, 2, 3},
			},
			want: 3,
		},
		{
			name: "2,2,1,1,1,2,2",
			args: args{
				nums: []int{2, 2, 1, 1, 1, 2, 2},
			},
			want: 2,
		},
		{
			name: "6,5,5",
			args: args{
				nums: []int{6, 5, 5},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MajorityElement(tt.args.nums), "MajorityElement(%v)", tt.args.nums)
		})
	}
}

func TestPascalTriangle(t *testing.T) {
	type args struct {
		rows int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "5",
			args: args{rows: 5},
			want: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			name: "1",
			args: args{rows: 1},
			want: [][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, PascalTriangle(tt.args.rows), "PascalTriangle(%v)", tt.args.rows)
		})
	}
}

func TestContainsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1,2,3,1",
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: true,
		},
		{
			name: "1,2,3,4",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: false,
		},
		{
			name: "1,1,1,3,3,4,3,2,4,2",
			args: args{
				nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ContainsDuplicate(tt.args.nums), "ContainsDuplicate(%v)", tt.args.nums)
		})
	}
}

func TestContainsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1,2,3,1;3",
			args: args{
				nums: []int{1, 2, 3, 1},
				k:    3,
			},
			want: true,
		},
		{
			name: "1,0,1,1",
			args: args{
				nums: []int{1, 0, 1, 1},
				k:    1,
			},
			want: true,
		},
		{
			name: "1,2,3,1,2,3",
			args: args{
				nums: []int{1, 2, 3, 1, 2, 3},
				k:    2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ContainsNearbyDuplicate(tt.args.nums, tt.args.k), "ContainsNearbyDuplicate(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}

func TestMinimumSize(t *testing.T) {
	type args struct {
		nums          []int
		maxOperations int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[9],2",
			args: args{
				nums:          []int{9},
				maxOperations: 2,
			},
			want: 3,
		},
		{
			name: "[2,4,8,2],4",
			args: args{
				nums:          []int{2, 4, 8, 2},
				maxOperations: 4,
			},
			want: 2,
		},
		{
			name: "[7,17],2",
			args: args{
				nums:          []int{7, 17},
				maxOperations: 2,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MinimumSize(tt.args.nums, tt.args.maxOperations), "MinimumSize(%v, %v)", tt.args.nums, tt.args.maxOperations)
		})
	}
}

func TestAlertNames(t *testing.T) {
	type args struct {
		keyName []string
		keyTime []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case1",
			args: args{
				keyName: []string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"},
				keyTime: []string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"},
			},
			want: []string{"daniel"},
		},
		{
			name: "case2",
			args: args{
				keyName: []string{"alice", "alice", "alice", "bob", "bob", "bob", "bob"},
				keyTime: []string{"12:01", "12:00", "18:00", "21:00", "21:20", "21:30", "23:00"},
			},
			want: []string{"bob"},
		},
		{
			name: "case3",
			args: args{
				keyName: []string{"john", "john", "john"},
				keyTime: []string{"23:58", "23:59", "00:01"},
			},
			want: []string{},
		},
		{
			name: "case4",
			args: args{
				keyName: []string{"leslie", "leslie", "leslie", "clare", "clare", "clare", "clare"},
				keyTime: []string{"13:00", "13:20", "14:00", "18:00", "18:51", "19:30", "19:49"},
			},
			want: []string{"clare", "leslie"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, AlertNames(tt.args.keyName, tt.args.keyTime), "AlertNames(%v, %v)", tt.args.keyName, tt.args.keyTime)
		})
	}
}

func TestRemoveSubfolders(t *testing.T) {
	type args struct {
		folder []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case1",
			args: args{
				folder: []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"},
			},
			want: []string{"/a", "/c/d", "/c/f"},
		},
		{
			name: "case2",
			args: args{
				folder: []string{"/a", "/a/b/c", "/a/b/d"},
			},
			want: []string{"/a"},
		},
		{
			name: "case3",
			args: args{
				folder: []string{"/a/b/c", "/a/b/ca", "/a/b/d"},
			},
			want: []string{"/a/b/c", "/a/b/ca", "/a/b/d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, RemoveSubfolders(tt.args.folder), "RemoveSubfolders(%v)", tt.args.folder)
		})
	}
}

func TestFourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case1",
			args: args{
				nums:   []int{1, 0, -1, 0, -2, 2},
				target: 0,
			},
			want: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			name: "case2",
			args: args{
				nums:   []int{2, 2, 2, 2, 2},
				target: 8,
			},
			want: [][]int{{2, 2, 2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, FourSum(tt.args.nums, tt.args.target), "FourSum(%v, %v)", tt.args.nums, tt.args.target)
		})
	}
}

func TestLongestWPI(t *testing.T) {
	type args struct {
		hours []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{hours: []int{6, 6, 6}},
			want: 0,
		},
		{
			name: "case2",
			args: args{hours: []int{9, 9, 6, 0, 6, 6, 9}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, LongestWPI(tt.args.hours), "LongestWPI(%v)", tt.args.hours)
		})
	}
}

func TestNumberOfPairs(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{nums: []int{1, 3, 2, 1, 3, 2, 2}},
			want: []int{3, 1},
		},
		{
			name: "case2",
			args: args{nums: []int{1, 1}},
			want: []int{1, 0},
		},
		{
			name: "case3",
			args: args{nums: []int{0}},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NumberOfPairs(tt.args.nums), "NumberOfPairs(%v)", tt.args.nums)
		})
	}
}

func TestGenerateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case1",
			args: args{n: 1},
			want: []string{"()"},
		},
		{
			name: "case2",
			args: args{n: 2},
			want: []string{"(())", "()()"},
		},
		{
			name: "case3",
			args: args{n: 3},
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GenerateParenthesis(tt.args.n), "GenerateParenthesis(%v)", tt.args.n)
		})
	}
}

func TestPermute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case1",
			args: args{[]int{1, 2, 3}},
			want: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		{
			name: "case2",
			args: args{nums: []int{1}},
			want: [][]int{{1}},
		},
		{
			name: "case3",
			args: args{nums: []int{1, 0}},
			want: [][]int{{1, 0}, {0, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Permute(tt.args.nums), "Permute(%v)", tt.args.nums)
		})
	}
}

func TestSearchRange(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				arr:    []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			want: []int{3, 4},
		},
		{
			name: "case2",
			args: args{
				arr:    []int{5, 7, 7, 8, 8, 10},
				target: 6,
			},
			want: []int{-1, -1},
		},
		{
			name: "case3",
			args: args{
				arr:    []int{},
				target: 0,
			},
			want: []int{-1, -1},
		},
		{
			name: "case4",
			args: args{
				arr:    []int{1},
				target: 1,
			},
			want: []int{0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SearchRange(tt.args.arr, tt.args.target), "SearchRange(%v, %v)", tt.args.arr, tt.args.target)
		})
	}
}

func TestMovesToMakeZigzag(t *testing.T) {
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
			args: args{nums: []int{1, 2, 3}},
			want: 2,
		},
		{
			name: "case2",
			args: args{nums: []int{9, 6, 1, 6, 2}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MovesToMakeZigzag(tt.args.nums), "MovesToMakeZigzag(%v)", tt.args.nums)
		})
	}
}

func TestMostFrequentEven(t *testing.T) {
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
			args: args{
				nums: []int{29, 47, 21, 41, 13, 37, 25, 7},
			},
			want: -1,
		},
		{
			name: "case2",
			args: args{
				nums: []int{0, 1, 2, 2, 4, 4, 1},
			},
			want: 2,
		},
		{
			name: "case3",
			args: args{
				nums: []int{4, 4, 4, 9, 2, 4},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MostFrequentEven(tt.args.nums), "MostFrequentEven(%v)", tt.args.nums)
		})
	}
}

func TestMergeOverlapping(t *testing.T) {

	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{
			name:      "case1",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want:      [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "case2",
			intervals: [][]int{{1, 4}, {4, 5}},
			want:      [][]int{{1, 5}},
		},
		{
			name:      "case3",
			intervals: [][]int{{1, 2}},
			want:      [][]int{{1, 2}},
		},
		{
			name:      "case4",
			intervals: [][]int{{1, 4}, {0, 4}},
			want:      [][]int{{0, 4}},
		},
		{
			name:      "case5",
			intervals: [][]int{{1, 4}, {2, 3}},
			want:      [][]int{{1, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MergeOverlapping(tt.intervals), "MergeOverlapping(%v)", tt.intervals)
		})
	}
}

func TestEliminateMaximum(t *testing.T) {
	type args struct {
		dist  []int
		speed []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				dist:  []int{1, 3, 4},
				speed: []int{1, 1, 1},
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				dist:  []int{1, 1, 2, 3},
				speed: []int{1, 1, 1, 1},
			},
			want: 1,
		},
		{
			name: "case3",
			args: args{
				dist:  []int{3, 2, 4},
				speed: []int{5, 3, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, EliminateMaximum(tt.args.dist, tt.args.speed), "EliminateMaximum(%v, %v)", tt.args.dist, tt.args.speed)
		})
	}
}

func TestNextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: []int{1, 3, 2},
		},
		{
			name: "case2",
			args: args{
				nums: []int{1, 1, 5},
			},
			want: []int{1, 5, 1},
		},
		{
			name: "case3",
			args: args{
				nums: []int{3, 2, 1},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NextPermutation(tt.args.nums)
			assert.Equalf(t, tt.want, tt.args.nums, "NextPermutation(%v)", tt.args.nums)
		})
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
