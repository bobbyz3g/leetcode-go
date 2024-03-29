package algebra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsUgly(t *testing.T) {
	var tests263 = []struct {
		n  int
		ok bool
	}{
		{6, true},
		{8, true},
		{14, false},
		{1, true},
		{0, false},
	}

	for _, v := range tests263 {
		assert.Equal(t, v.ok, IsUgly(v.n))
	}
}

func TestGetRow(t *testing.T) {
	var tests119 = []struct {
		a int
		b []int
	}{
		{
			a: 1,
			b: []int{1, 1},
		},
		{
			a: 3,
			b: []int{1, 3, 3, 1},
		},
	}
	for _, v := range tests119 {
		assert.Equal(t, GetRow(v.a), v.b)
	}
}

func TestFindDisappearedNumbers(t *testing.T) {
	var tests448 = []struct {
		a []int
		b []int
	}{
		{
			a: []int{4, 3, 2, 7, 8, 2, 3, 1},
			b: []int{5, 6},
		},
	}
	for _, v := range tests448 {
		assert.Equal(t, FindDisappearedNumbers(v.a), v.b)
	}
}

func TestIsValidSerialization(t *testing.T) {

	var tests331 = []struct {
		preorder string
		ok       bool
	}{
		{
			"9,3,4,#,#,1,#,#,2,#,6,#,#",
			true,
		},
		{
			"1,#",
			false,
		},
		{
			"9,#,#,1",
			false,
		},
		{
			"#",
			true,
		},
	}
	for _, v := range tests331 {
		assert.Equal(t, v.ok, IsValidSerialization(v.preorder))
	}
}

func TestLongestOnes(t *testing.T) {
	var tests1004 = []struct {
		a      []int
		k      int
		output int
	}{
		{
			a:      []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			k:      2,
			output: 6,
		},
		{
			a:      []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			k:      3,
			output: 10,
		},
	}
	for _, v := range tests1004 {
		assert.Equal(t, LongestOnes(v.a, v.k), v.output)
	}
}

func TestMinKBitFlips(t *testing.T) {

	var tests35 = []struct {
		nums   []int
		target int
		output int
	}{
		{
			nums:   []int{1, 3, 5, 6},
			target: 5,
			output: 2,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 2,
			output: 1,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 7,
			output: 4,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 0,
			output: 0,
		},
	}
	for _, v := range tests35 {
		assert.Equal(t, SearchInsert(v.nums, v.target), v.output)
	}
}

func TestIToR(t *testing.T) {
	var tests12 = []struct {
		num int
		r   string
		err error
	}{
		{
			3,
			"III",
			nil,
		},
		{
			4,
			"IV",
			nil,
		},
		{
			9,
			"IX",
			nil,
		},
		{
			58,
			"LVIII",
			nil,
		},
		{
			1994,
			"MCMXCIV",
			nil,
		},
		{
			0,
			"",
			ErrOutOfBound,
		},
	}

	for _, tt := range tests12 {
		r, err := IToR(tt.num)
		assert.Equal(t, tt.r, r)
		assert.Equal(t, tt.err, err)
	}
}

func TestMaximumUnits(t *testing.T) {
	testcases := []struct {
		boxes    [][]int
		size     int
		maxUnits int
	}{
		{boxes: [][]int{{1, 3}, {2, 2}, {3, 1}}, size: 4, maxUnits: 8},
		{boxes: [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, size: 10, maxUnits: 91},
		{boxes: [][]int{{1, 3}, {2, 2}, {3, 1}}, size: 20, maxUnits: 10},
	}

	for _, tc := range testcases {
		assert.Equal(t, tc.maxUnits, MaximumUnits(tc.boxes, tc.size))
	}
}

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		want   []int
	}{
		{"one", []int{1}, []int{2}},
		{"nine", []int{9}, []int{1, 0}},
		{"ten", []int{1, 0}, []int{1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, PlusOne(tt.digits), "PlusOne(%v)", tt.digits)
		})
	}
}

func TestSqlX(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{x: 1}, want: 1},
		{name: "2", args: args{x: 2}, want: 1},
		{name: "4", args: args{x: 4}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SqrtX(tt.args.x), "SqrtX(%v)", tt.args.x)
		})
	}
}

func TestClimbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{n: 1}, want: 1},
		{name: "2", args: args{n: 2}, want: 2},
		{name: "3", args: args{n: 3}, want: 3},
		{name: "4", args: args{n: 4}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ClimbStairs(tt.args.n), "ClimbStairs(%v)", tt.args.n)
		})
	}
}

func TestIsHappyNum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{name: "1", args: args{n: 1}, want: true},
		{name: "2", args: args{n: 2}, want: false},
		{name: "19", args: args{n: 19}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsHappyNum(tt.args.n), "IsHappyNum(%v)", tt.args.n)
		})
	}
}

func TestMaximumScore(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2,4,6",
			args: args{
				a: 2,
				b: 4,
				c: 6,
			},
			want: 6,
		},
		{
			name: "4,4,6",
			args: args{
				a: 4,
				b: 4,
				c: 6,
			},
			want: 7,
		},
		{
			name: "1,8,8",
			args: args{
				a: 1,
				b: 8,
				c: 8,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MaximumScore(tt.args.a, tt.args.b, tt.args.c), "MaximumScore(%v, %v, %v)", tt.args.a, tt.args.b, tt.args.c)
		})
	}
}

func TestReinitializePermutation(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "4",
			args: args{n: 4},
			want: 2,
		},
		{
			name: "6",
			args: args{n: 6},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ReinitializePermutation(tt.args.n), "ReinitializePermutation(%v)", tt.args.n)
		})
	}
}

func TestSmallestRepunitDivByK(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{k: 1},
			want: 1,
		},
		{
			name: "case2",
			args: args{k: 2},
			want: -1,
		},
		{
			name: "case3",
			args: args{k: 3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SmallestRepunitDivByK(tt.args.k), "SmallestRepunitDivByK(%v)", tt.args.k)
		})
	}
}

func TestThreeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				nums:   []int{-1, 1, -4, 2},
				target: 1,
			},
			want: 2,
		},
		{
			name: "case2",
			args: args{
				nums:   []int{0, 0, 0},
				target: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ThreeSumClosest(tt.args.nums, tt.args.target), "ThreeSumClosest(%v, %v)", tt.args.nums, tt.args.target)
		})
	}
}

func TestSplitNum(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{num: 4325},
			want: 59,
		},
		{
			name: "case2",
			args: args{num: 687},
			want: 75,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SplitNum(tt.args.num), "SplitNum(%v)", tt.args.num)
		})
	}
}

func TestSumIndicesWithKSetBits(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{5, 10, 1, 5, 2},
				k:    1,
			},
			want: 13,
		},
		{
			name: "case2",
			args: args{
				nums: []int{4, 3, 2, 1},
				k:    2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SumIndicesWithKSetBits(tt.args.nums, tt.args.k), "SumIndicesWithKSetBits(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
